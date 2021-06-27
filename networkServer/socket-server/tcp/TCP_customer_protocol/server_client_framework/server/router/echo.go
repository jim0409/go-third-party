package router

import (
	"encoding/json"
	"log"
)

type EchoController struct{}

func (e *EchoController) Excute(m Message) []byte {
	log.Println("Receive the msg ", m)

	m.Meta["echo"] = "ack"
	msg, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return nil
	}

	return msg
}

func init() {
	var echo EchoController

	Route(func(m Message) bool { // Route註冊echo這個資料結構到router上
		return m.Meta["meta"] == "pass"
	}, &echo)
}
