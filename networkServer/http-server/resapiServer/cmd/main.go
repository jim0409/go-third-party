package main

import (
	"io"
	"log"
	"net/http"

	"github.com/jimweng/networkServer/http-server/resapiServer/behavior/sayhi"
)

// Descriptions of Operation steps:
// open terminial and run this program
// curl localhost:9000/${EPT}
// EPT : depends on end point would return corresponding response
// e.g.
// terminal1$: go run main.go
// terminal2$: curl localhost:9000/queryURL ; output should be 'the url is /queryURL'

type myHandler int

func (h myHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	switch req.URL.Path {
	case "/cat":
		io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/0/06/Kitten_in_Rizal_Park%2C_Manila.jpg">`)
	case "/dog":
		io.WriteString(res, `<img src="https://upload.wikimedia.org/wikipedia/commons/6/6e/Golde33443.jpg">`)
	case "/sayhi":
		var s sayhi.SayHi
		s.Name = "jim"
		s.Birthday = "April 9"
		io.WriteString(res, s.HI())
	case "/queryURL":
		var s sayhi.SayHi
		io.WriteString(res, s.GetURL(req))
	default:
		io.WriteString(res, "No such api")
	}
}

func main() {
	log.Printf("you can use curl or browser to interact with your server now!")

	var h myHandler
	http.ListenAndServe(":9000", h)
}
