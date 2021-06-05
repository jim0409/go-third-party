package router

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PassBody struct {
	Param interface{}
}

type RequestBody struct {
	Param interface{}
	Sign  string
}

// UserMap為假定，合法的使用者
var UserMap = map[string]int{"jim": 1, "tim": 1}

// AuthCheck是一個前置檢查動作，所有請求都需要滿足的
func AuthCheck() gin.HandlerFunc {
	f := func(c *gin.Context) {
		user := c.GetHeader("User")
		switch UserMap[user] {
		case 1:
			return
		default:
			invalid(c)

		}
	}
	return f
}

// 非法使用者的handlerFunc
func invalid(c *gin.Context) {
	c.JSON(400, gin.H{
		"code":   "2",
		"result": "failed",
		"msg":    "Invalid User.",
	})
	c.Abort()
	return
}

// BodyPreProcessing 使用要注意，他是全部的body做解密
func BodyPreProcessing() gin.HandlerFunc {
	f := func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Context() != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		decryptMsg, err := base64.StdEncoding.DecodeString(string(bodyBytes))
		if err != nil {
			c.Next()
		}

		log.Printf("After base64 -D request body is %v\n", string(decryptMsg))

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(decryptMsg))
	}
	return f
}

func AddLog() gin.HandlerFunc {
	f := func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Context() != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		log.Printf("Receive Raw Request body with %v\n", string(bodyBytes))

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
	}
	return f
}

func ValidateSignature() gin.HandlerFunc {
	f := func(c *gin.Context) {
		var bodyBytes []byte
		if c.Request.Context() != nil {
			bodyBytes, _ = ioutil.ReadAll(c.Request.Body)
		}

		// m := make(map[string]interface{})
		m := &RequestBody{}
		if err := json.Unmarshal(bodyBytes, m); err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid to do json unmarshal with err: %v", err))
			c.Abort()
			return
		}
		if m.Sign != "good" {
			c.JSON(http.StatusBadRequest, "SignError")
			c.Abort()
			return
		}

		passMap := &PassBody{}
		passMap.Param = m.Param

		newbodyBytes, err := json.Marshal(passMap)
		if err != nil {
			c.JSON(http.StatusBadRequest, fmt.Sprintf("Invalid to do json marshal with err: %v", err))
			c.Abort()
			return
		}

		log.Printf("After validate Request body becomes %v\n", string(newbodyBytes))

		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(newbodyBytes))
	}
	return f
}

// CheckParamAndHeader是一個裝飾器方法，透過傳入的`Context`回傳對應的Handler
func CheckParamAndHeader(h gin.HandlerFunc) gin.HandlerFunc {
	f := func(c *gin.Context) {
		header := c.Request.Header.Get("token")
		if header == "" {
			c.JSON(400, gin.H{
				"code":   "3",
				"result": "failed",
				"msg":    "Missing token",
			})
			return
		} else {
			h(c)
		}
	}

	return f

}
