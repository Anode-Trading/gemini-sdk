package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func StreamOrderBook(pair string, stopChan chan struct{}, geminiClosedChan chan<- struct{}, dataChan chan OrderBookResponse) {

	url := fmt.Sprint(baseUrl, v1, orderBookPath, pair)
	c, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {
			fmt.Println("close connection err:", err)
		}
	}(c)

	lastSequenceNumber := -1
	doneStreaming := false
	for !doneStreaming {
		select {
		case <-stopChan:
			log.Println("closing connection as requested")
			err := c.WriteMessage(websocket.CloseMessage, []byte{})
			if err != nil {
				log.Println("close connection err:", err)
			}
			doneStreaming = true
		default:
			_, message, err := c.ReadMessage()
			if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Println("connection closed by gemini server:", err)
				geminiClosedChan <- struct{}{}
				close(geminiClosedChan)
				close(dataChan)
				doneStreaming = true
				continue
			}

			if err != nil {
				log.Println("err read:", err)
				geminiClosedChan <- struct{}{}
				close(geminiClosedChan)
				close(dataChan)
				doneStreaming = true
				continue
			}
			var orderBookResponse OrderBookResponse
			err = json.Unmarshal(message, &orderBookResponse)
			if err != nil {
				fmt.Println("err unmarshal:", err)
			}
			if lastSequenceNumber+1 != orderBookResponse.SocketSequenceNumber {
				log.Println("incorrect sequence, try again", orderBookResponse.SocketSequenceNumber)
				geminiClosedChan <- struct{}{}
				close(geminiClosedChan)
				close(dataChan)
				doneStreaming = true
				continue
			}
			select {
			case <-stopChan:
				log.Println("closing connection as requested")
				err := c.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Println("close connection err:", err)
				}
				doneStreaming = true
			default:
				dataChan <- orderBookResponse
				lastSequenceNumber++
			}

		}
	}
	log.Println("orderBook streaming stopped by gemini server")
}
