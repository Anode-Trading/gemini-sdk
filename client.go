package main

import (
	"fmt"
	"github.com/Anode-Trading/gemini/api"
	"log"
	"time"
)

func clientOrderBook() {
	stopChan, doneStream := make(chan struct{}), make(chan struct{})
	dataChan := make(chan api.OrderBookResponse)
	go api.StreamOrderBook("BTCUSD", stopChan, doneStream, dataChan)

	timer := time.NewTimer(20 * time.Second)

	isDone := false
	for !isDone {
		select {
		case <-doneStream:
			log.Println("streaming stopped from gemini, hence closing receiver")
			isDone = true
		case res := <-dataChan:
			//Do your stuff
			log.Println("seq num : ", res.SocketSequenceNumber)
		case <-timer.C: // To notify that you are no longer interested in receiving updates
			select {
			case res := <-dataChan:
				// Essential to stop deadlock
				// Do your stuff
				log.Println("seq num : ", res.SocketSequenceNumber)
			}
			stopChan <- struct{}{}
			close(stopChan)
			isDone = true
		}
	}
	log.Println("shutting down")

}

func clientSecurities() {
	securities, err := api.GetSecurities()
	if err != nil {
		return
	}
	fmt.Println(securities)
}

func clientSecurityInfo() {
	info, err := api.GetSecurityInfo("btcusd")
	if err != nil {
		return
	}
	fmt.Println(info)
}

func main() {
	clientOrderBook()
}
