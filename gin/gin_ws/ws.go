package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func WebsocketHandler(c *gin.Context) {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		// TODO: ...
		panic(err)
	}

	defer func() {
		closeSocketErr := ws.Close()
		if closeSocketErr != nil {
			// TODO: ...
			panic(err)
		}
	}()

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			// TODO: ...
			panic(err)
		}

		// TODO: ...
		fmt.Printf("Message Type: %d, Message: %s\n", msgType, string(msg))

		err = ws.WriteJSON(struct {
			Reply string `json:"reply"`
		}{
			Reply: fmt.Sprintf("Echo ... %v", string(msg)),
		})

		if err != nil {
			// TODO: ...
			panic(err)
		}
	}
}
