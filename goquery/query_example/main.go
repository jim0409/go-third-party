package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

var queryUrl = "https://github.com/login"

func getElementById(id string, n *html.Node) (element *html.Node, ok bool) {
	for _, a := range n.Attr {
		if a.Key == "id" && a.Val == id {
			return n, true
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if element, ok = getElementById(id, c); ok {
			return
		}
	}
	return
}

func getKeyValue() {
	resp, err := http.Get(queryUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	root, err := html.Parse(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	element, ok := getElementById("login_field", root)
	if !ok {
		log.Fatal("element not found")
	}

	for i, j := range element.Attr {
		fmt.Printf("%v___%v\n", i, j)
	}

}

func main() {
	// refer : https://www.reddit.com/r/golang/comments/3fcabt/question_read_value_from_html_input_tag/
	getKeyValue()
}
