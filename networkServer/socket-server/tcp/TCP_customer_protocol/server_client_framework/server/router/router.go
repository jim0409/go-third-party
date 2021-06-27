package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Msg struct {
	Meta    map[string]interface{} `json:"meta"`
	Content interface{}            `json:"content"`
}

type Controller interface {
	Excute(Msg) []byte
}

var routers = make([][2]interface{}, 0)

// Router : 透過解析pred的類別，來將controller正確定義到 `routers` 上
func Route(pred interface{}, controller Controller) {
	switch pred.(type) {
	/*
		如果傳入的 pred 是一個函數 input: `Msg` return `bool`，直接進行append
	*/
	case func(entry Msg) bool:
		{
			var arr [2]interface{}
			arr[0] = pred
			arr[1] = controller
			routers = append(routers, arr)
		}
	/*
		如果傳入的 pred 是一個字典，定義一個函數defaultPred符合格式func(entry Msg) bool，在進行append
	*/
	case map[string]interface{}:
		// defaultPred 定義了一個封閉函數
		defaultPred := func(entry Msg) bool {
			for keyPred, valPred := range pred.(map[string]interface{}) {
				val, ok := entry.Meta[keyPred]
				if !ok {
					return false
				}
				if val != valPred {
					return false
				}
			}
			return true
		}

		routers = append(routers, [2]interface{}{defaultPred, controller})

		fmt.Println(routers)
	default:
		fmt.Println("No match requested controller")
	}
}

func TaskDeliver(postdata []byte, conn net.Conn) {
	for _, v := range routers {
		pred := v[0]
		act := v[1]
		var entermsg Msg
		err := json.Unmarshal(postdata, &entermsg)
		if err != nil {
			log.Println(err)
		}

		if pred.(func(entermsg Msg) bool)(entermsg) {
			result := act.(Controller).Excute(entermsg)

			conn.Write(result)
			return
		}

		conn.Write([]byte("Hey! The meta data is invalid!"))
		return
	}
}

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

	// Route註冊echo這個資料結構到router上
	Route(func(entry Msg) bool {
		return entry.Meta["meta"] == "pass"
	}, &echo)
}
