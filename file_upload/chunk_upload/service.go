package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	r := c.Request
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf("Error Retrieving the File %v\n", err)
		return
	}

	defer file.Close()

	id, err := BackUpFile(file, handler.Filename)
	if err != nil {
		c.JSON(404, fmt.Sprintf("Failed to Uploaded File %v", err))
		return
	}

	c.JSON(200, gin.H{
		"status":   "success",
		"id":       id,
		"filename": fmt.Sprintf("Uploaded File: %+v", handler.Filename),
		"size":     fmt.Sprintf("File Size: %+v", handler.Size),
		"header":   fmt.Sprintf("MIME Header: %+v", handler.Header),
	})
}

func BackUpFile(file io.Reader, filename string) (string, error) {
	if _, err := os.Stat("files"); os.IsNotExist(err) {
		if err := os.Mkdir("files", 0777); err != nil {
			return "", err
		}
	}
	tempFile, err := ioutil.TempFile("files", fmt.Sprintf("upload-*-%s.backup", filename)) // `*` 會隨機產生一個亂序 id
	if err != nil {
		return "", err
	}
	defer tempFile.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return "", err
	}

	// can't use tempFile since in memory bytes differ from file md5
	// TODO:  tempFile.Name() should record into db.
	id, err := LoadFileMD5(tempFile.Name())
	if err != nil {
		return "", err
	}

	return id, nil
}

func LoadFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	return FileMD5(file)
}

func FileMD5(file *os.File) (string, error) {
	hash := md5.New()
	_, _ = io.Copy(hash, file)
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// 分片校驗
func CheckUploadFile(c *gin.Context) {
	// 1. query info from db
	// 2. according to file md5 return status; 0: not existed, 1: uploading, 2: existed, 99: md5 value err
}
