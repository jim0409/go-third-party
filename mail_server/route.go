package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func apiRouter(route *gin.Engine) {

	group := route.Group("/")
	group.GET("_health", func(c *gin.Context) { c.JSON(200, gin.H{"msg": "working env: " + GlobalEnv}) })

	msg := group.Group("/msg")
	{
		msg.Use(MiddleWare())
		msg.POST("/send", SendHandler)
	}

}

type PostBody struct {
	AuthCode string `json:"authcode"` // for smtp server authentication? or middleware check?
	ID       int    `json:"id"`
	Mail     string `json:"mail"`
	Sub      string `json:"sub"`
	From     string `json:"from"`
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		var bs []byte
		var err error
		if c.Request.Body == nil {
			c.JSON(400, gin.H{"msg": "post body required!"})
			c.Abort()
			return
		}

		bs, err = ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(400, err)
			c.Abort()
			return
		}
		fmt.Println(string(bs))

		postbd := &PostBody{}
		err = json.Unmarshal(bs, postbd)
		if err != nil {
			c.JSON(400, gin.H{"msg": err})
			c.Abort()
			return
		}

		if postbd.AuthCode == "" {
			c.JSON(400, gin.H{"msg": "lack of auth code!"})
			c.Abort()
			return
		}

		// according to verify rules to use authcode !!

		if postbd.ID == 0 {
			c.JSON(400, gin.H{"msg": "lack of template id!"})
			c.Abort()
			return
		}

		if postbd.Mail == "" {
			c.JSON(400, gin.H{"msg": "lack of sender mail!"})
			c.Abort()
			return
		}

		if postbd.Sub == "" {
			postbd.Sub = DemoSubject
		}

		if postbd.From == "" {
			postbd.From = DemoFrom
		}

		// Restore the io.ReadCloser to its original state
		// c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bs))
		c.Set("id", postbd.ID)
		c.Set("mail", postbd.Mail)
		c.Set("sub", postbd.Sub)
		c.Set("from", postbd.From)

		c.Next()
	}
}

func SendHandler(c *gin.Context) {
	id := c.GetInt("id")
	mail := c.GetString("mail")
	sub := c.GetString("sub")
	from := c.GetString("from")
	sender := NewSender(GlobalAuth, GlobalUser, GlobalHost, GlobalSmtpserver, &MsgTemplate)
	err := sender.SendMail(from, mail, sub, id)
	if err != nil {
		c.JSON(400, gin.H{"msg": err})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"msg": fmt.Sprintf("post id %v to mail %v", id, mail)})
}
