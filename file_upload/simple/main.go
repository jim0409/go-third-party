package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func BackUpFile(file io.Reader) error {
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		if err := os.Mkdir("files", 0777); err != nil {
			return err
		}
	}
	tempFile, err := ioutil.TempFile("files", "upload-*.backup") // `*` 會隨機產生一個亂序 id
	if err != nil {
		return err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	_, err = tempFile.Write(fileBytes)

	return err
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

		if err := BackUpFile(file); err != nil {
			fmt.Fprintf(w, "Failed to Uploaded File\n")
			fmt.Println(err)
		} else {
			fmt.Fprintf(w, "Successfully Uploaded File\n")
		}
	}
}

func setupRoutes() {
	http.HandleFunc("/upload", uploadFile)
	http.ListenAndServe(":2020", nil)
}

func main() {
	setupRoutes()
}
