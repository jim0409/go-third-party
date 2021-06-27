package router

import (
	"encoding/json"
	"log"
)

type EchoController struct{}

func (e *EchoController) Excute(m Msg) []byte {
	log.Println("Receive the msg ", m)

	m.Meta["echo"] = "ack"
	msg, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	return msg
}

func init() {
	var echo EchoController

	Route(func(entry Msg) bool { // Route註冊echo這個資料結構到router上
		return entry.Meta["meta"] == "pass"
	}, &echo)
}
