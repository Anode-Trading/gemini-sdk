package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
)

func StreamOrderBook(pair string, stopChan chan struct{}, dataChan chan OrderBookResponse) {

	url := fmt.Sprint(baseUrl, v1, orderBookPath, pair)
	fmt.Println(url)
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
	for {
		_, message, err := c.ReadMessage()
		if err != nil && websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
			log.Println("connection closed by gemini server:", err)
			stopChan <- struct{}{}
			close(dataChan)
			return
		}

		if err != nil {
			log.Println("err read:", err)
			stopChan <- struct{}{}
			close(dataChan)
			return
		}
		var orderBookResponse OrderBookResponse
		err = json.Unmarshal(message, &orderBookResponse)
		if err != nil {
			fmt.Println("err unmarshal:", err)
		}
		if lastSequenceNumber+1 != orderBookResponse.SocketSequenceNumber {
			log.Println("incorrect sequence", orderBookResponse.SocketSequenceNumber)
			stopChan <- struct{}{}
			break
		}

		dataChan <- orderBookResponse
		lastSequenceNumber++
	}
}
