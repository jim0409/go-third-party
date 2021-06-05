package echo

import (
	"encoding/json"
	"log"

	"github.com/jimweng/networkServer/socket-server/udp/UDP_customer_protocol/server/router"
	"github.com/jimweng/networkServer/socket-server/udp/UDP_customer_protocol/utils"
)

type EchoController struct{}

func (e *EchoController) Execute(m utils.Msg) []byte {
	log.Println("Receive the msg ", m)

	m.Meta["echo"] = "ack"
	msg, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	return msg
}

func init() {
	router.Route(func(entry utils.Msg) bool {
		if entry.Meta["meta"] == "test" { // if entry.Meta value is "test" return true, else return false
			return true
		}
		return false
	}, &EchoController{})
}
