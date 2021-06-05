package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	resp, err := http.Get("https://golang.google.cn/pkg/")
	if err != nil {
		// handle err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("Error response with status code: %v\n", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc, _ := html.Parse(bytes.NewReader(bodyBytes))

	bn, err := getBody(doc)
	if err != nil {
		return
	}
	body := renderNode(bn)
	// fmt.Printf("The type of body is %v\n", reflect.TypeOf(body))
	fmt.Println(body)
}

func getBody(doc *html.Node) (*html.Node, error) {
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "table" {
			// fmt.Printf("the wanted value is ... %v__%v\n", n.Data)
			// fmt.Printf("the wanted value is ... %v\n", n.LastChild.Data)
			// fmt.Printf("the wanted value is ... %v\n", n.FirstChild.Data)
			// fmt.Printf("the wanted value is ... %v\n", n.PrevSibling.Data)
			// fmt.Printf("the wanted value is ... %v\n", n.Parent.Data)
			// // fmt.Printf("the wanted value is ... %v\n", n.NextSibling.Data)

			// fmt.Printf("the wanted value is ... %v\n", n.Data) // body
			b = n
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func renderNode(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}
