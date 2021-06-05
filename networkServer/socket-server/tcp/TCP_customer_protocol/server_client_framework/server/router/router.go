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

// 定義一個routers，並且在該router註冊處理方法
var routers [][2]interface{}

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
		var arr [2]interface{}
		// 將defaultPred存入arr[0]
		arr[0] = defaultPred
		arr[1] = controller
		routers = append(routers, arr)
		fmt.Println(routers)
	default:
		fmt.Println("No match requested controller")
	}
}

// TaskDeliver 會輸入預計接收資料的bufferBytes以及conn(連線)
func TaskDeliver(postdata []byte, conn net.Conn) {
	for _, v := range routers {
		pred := v[0]
		act := v[1]
		var entermsg Msg
		err := json.Unmarshal(postdata, &entermsg)
		if err != nil {
			log.Println(err)
		}

		// pred這邊定義的函數為來自router的Controller，只有當Controller回傳為true的時候才進行寫入
		if pred.(func(entermsg Msg) bool)(entermsg) {
			result := act.(Controller).Excute(entermsg)

			// 把結果(result)寫回conn，期許client端也會收到對應的資訊
			conn.Write(result)
			return
		} else {
			conn.Write([]byte("Hey! The meta data is invalid!"))
			return
		}
	}
}

// EchoController 為一個controller的實例，所有controller都需要在init()時註冊到router，才能被router所分配
type EchoController struct {
}

func (e *EchoController) Excute(m Msg) []byte {
	log.Println("Receive the msg ", m)

	m.Meta["echo"] = "ack"
	msg, err := json.Marshal(m)
	if err != nil {
		return nil
	}

	return msg
}

// 透過初始化設定，將已定義的Execute方法的`資料結構`註冊到router
func init() {
	var echo EchoController

	// 實例routers
	routers = make([][2]interface{}, 0, 20)

	// Route註冊echo這個資料結構到router上
	Route(func(entry Msg) bool {
		// 如果entry.Meta的值為"test"則返回true
		if entry.Meta["meta"] == "test" {
			return true
		}
		return false
	}, &echo)
}
