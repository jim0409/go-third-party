package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-third-party/gin/customerized/controllers/apps"
	. "go-third-party/gin/customerized/utils"
)

type RequestBody struct {
	Param interface{}
}

func Collector(c *gin.Context) {
	reqBody, err := c.GetRawData()
	if err != nil {
		c.Abort()
	}

	// log.Printf("Inside controller received body is %v", string(reqBody))

	// claim GetResponse as doFunc
	recvReq := &RequestBody{}
	json.Unmarshal(reqBody, recvReq)

	workTask := WorkRequest{
		Execute:      apps.GetResponse,
		ExecuteInput: recvReq.Param,          // 使用 ExecuteInput 將傳入的body帶進WorkerRequest裡面，由於承接的 WorkRequest.ExecuteInput 是一個interface{}，所以可以再斟酌傳入的資料結構
		Result:       make(chan interface{}), // 如果不生成實例，會導致阻塞
	}

	// retrive worker from worker pools
	Wspools.RetriveWorkerQueue() <- workTask

	resp := <-workTask.Result

	// c.JSON(http.StatusOK, fmt.Sprintf("%v", resp))
	WebResp(c, http.StatusOK, noError, fmt.Sprintf("%v", resp))
	return
}

func WebResp(c *gin.Context, statusCode int, Message string, data interface{}) {
	respMap := map[string]interface{}{
		"Code":    statusCode,
		"Message": Message,
		"Data": map[string]interface{}{
			"content": data,
		},
	}
	c.JSON(statusCode, respMap)
	return
}

const (
	noError = "everything allright"
)
