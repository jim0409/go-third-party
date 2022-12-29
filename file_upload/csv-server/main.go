package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func ShowCSVRecord(file io.Reader) {
	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(record)
	}
}

func BackUpFile(file io.Reader) {
	tempFile, err := ioutil.TempFile("files", "upload-*.backup") // `*` 會隨機產生一個亂序 id
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	tempFile.Write(fileBytes)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("fileUp.gtpl")
		t.Execute(w, nil)
	} else {
		// 1. parse input
		r.ParseMultipartForm(10 << 20)
		// 2. retrieve file
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			fmt.Println("Error Retrieving the File")
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)

		// 3. write temporary file on our server
		// 因為 io.Reader 讀取完畢就會回收。所以一次只能開啟一個，不然就是要另外做 bytes.Buffer
		ShowCSVRecord(file)
		// BackUpFile(file)

		// 4. return result
		fmt.Fprintf(w, "Successfully Uploaded File\n")
	}
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":2020", nil)
}

func main() {
	setupRoutes()
}
