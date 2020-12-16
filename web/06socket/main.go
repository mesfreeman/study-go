package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// Echo 打印接收字符
func Echo(ws *websocket.Conn) {
	for {
		var reply string
		if err := websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Can`t receive")
			break
		}

		fmt.Println("Receive back from client: ", reply)
		msg := "Received: " + reply
		fmt.Println("Sending to client: ", msg)

		if err := websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Can`t send")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
