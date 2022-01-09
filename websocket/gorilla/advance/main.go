package main

import (
	"go-third-party/websocket/gorilla/advance/room"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/websocket"
)

var websocketManager room.WebSocketManagerIface

var GroupMap = map[string]map[string]bool{
	"group": map[string]bool{
		"jim": true,
	},
}

func dummyfindHdr(group string, usr string) bool {
	return GroupMap[group][usr]
}

func main() {

	webSocket := room.InitManager(dummyfindHdr)
	websocketManager = webSocket
	webSocket.HubStart()

	defer webSocket.HubStop()

	srv := &http.Server{
		Addr:    ":8000",
		Handler: http.HandlerFunc(connectWebSocket),
	}

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Printf("failed to start server %v\n", err)
		}
		log.Printf("server started listen at %v\n", err.Error())
	}()
	SendWebsocketMsg()

	interruptChan := make(chan os.Signal, 1)

	signal.Notify(interruptChan, syscall.SIGINT, syscall.SIGTERM)
	<-interruptChan

}

func connectWebSocket(w http.ResponseWriter, r *http.Request) {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  room.MaxMessageSize,
		WriteBufferSize: room.MaxMessageSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("api.web_socket.connect.upgrade.app_err: %s", err.Error())
		return
	}

	wc := websocketManager.NewWebSocketConn(ws, room.Session{
		UserID: "jim",
	})
	websocketManager.HubRegister(wc)

	log.Printf("%v ws connected", wc.UserID)

	wc.Pump()
}

// 定時向`group`發送訊息
func SendWebsocketMsg() {

	wse := room.NewWebSocketEvent(room.WebsocketEventMessage, "group", "", nil)
	wse.Data = "some event is comming soon ~"

	go func() {
		for {
			time.Sleep(1 * time.Second)
			websocketManager.Publish(wse)
		}
	}()
}
