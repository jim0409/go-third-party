package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/jimweng/networkServer/socket-server/udp/UDP_customer_protocol/utils"
)

func init() {
	Routers = make([][2]interface{}, 0, 20)
}

type Controller interface {
	Execute(utils.Msg) []byte
}

var Routers = [][2]interface{}{}

// Router : 透過解析pred的類別，來將controller正確定義到 `routers` 上
func Route(pred interface{}, controller Controller) {
	switch pred.(type) {
	/*
		如果傳入的 pred 是一個函數 input: `Msg` return `bool`，直接進行append
	*/
	case func(entry utils.Msg) bool:
		{
			var arr [2]interface{}
			arr[0] = pred
			arr[1] = controller
			Routers = append(Routers, arr)
		}
	/*
		如果傳入的 pred 是一個字典，定義一個函數defaultPred符合格式func(entry Msg) bool，在進行append
	*/
	case map[string]interface{}:
		// defaultPred 定義了一個封閉函數
		defaultPred := func(entry utils.Msg) bool {
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
		Routers = append(Routers, arr)
		fmt.Println(Routers)
	default:
		fmt.Println("No match requested controller")
	}
}

/*
	TaskDeliver would delive task to correspond methods
*/
func TaskDeliver(postdata []byte, listener net.PacketConn, addr net.Addr) {
	for _, v := range Routers {
		pred := v[0]
		act := v[1]
		var entermsg utils.Msg
		err := json.Unmarshal(postdata, &entermsg)
		if err != nil {
			log.Println(err)
		}

		if pred.(func(entermsg utils.Msg) bool)(entermsg) {
			result := act.(Controller).Execute(entermsg)
			listener.WriteTo(result, addr)
			return

		} else {
			continue // can't use return or router would not work through
		}
	}
}

// TODO: add Worker Pool instead of for loop to check TaskDeliver
