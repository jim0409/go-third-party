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

	// TODO: retrieve user from JWT
	// TODO: retrieve chunk num from query params
	// TODO: retrieve request md5 from query params, check md5 before save another chunk file

	// md5value := "9176b139835b4888ef37776bfdeefab6"
	md5value := c.Query("md5value")
	if md5value == "" {
		c.JSON(400, "lack of md5value!")
		return
	}

	// filename := "docker-compose.yml"
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(400, "lack of filename!")
		return
	}

	// username := "jim"
	username := c.GetHeader("username")
	if username == "" {
		c.JSON(400, "lack of username!")
		return
	}

	uploadfile, err := opdb.FindUploadDetailByFileName(md5value, filename)
	if err != nil {
		c.JSON(404, fmt.Sprintf("Failed to Uploaded File %v", err))
		return
	}
	if uploadfile.ID != 0 {
		c.JSON(200, gin.H{
			"status":   "success",
			"id":       uploadfile.ID,
			"filename": fmt.Sprintf("Uploaded File: %+v", uploadfile.FileName),
			"size":     fmt.Sprintf("File Size: %+v", uploadfile.Size),
		})
		return
	}

	// TODO: if size over restrict, return upload error msg
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf("Error Retrieving the File %v\n", err)
		return
	}
	defer file.Close()

	size := handler.Size
	id, err := BackUpFile(file, username, handler.Filename, md5value, size, 1)
	if err != nil {
		c.JSON(404, fmt.Sprintf("Failed to Uploaded File %v", err))
		return
	}

	c.JSON(200, gin.H{
		"status":   "success",
		"id":       id,
		"filename": fmt.Sprintf("Uploaded File: %+v", handler.Filename),
		"size":     fmt.Sprintf("File Size: %+v", handler.Size),
		// "header":   fmt.Sprintf("MIME Header: %+v", handler.Header),
	})
}

func BackUpFile(file io.Reader, usrname string, filename string, md5value string, size int64, chunknum int) (int, error) {

	id, err := opdb.InsertOneRecord(usrname, filename, md5value, size, chunknum)
	if err != nil {
		return 0, err
	}

	// after check file status, not exsited and save another tmpFile
	tempFile, err := ioutil.TempFile("files", fmt.Sprintf("upload-*-%s.backup", filename)) // `*` 會隨機產生一個亂序 id
	if err != nil {
		return 0, err
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, err
	}

	_, err = tempFile.Write(fileBytes)
	if err != nil {
		return 0, err
	}

	// can't use tempFile since in memory bytes differ from file md5
	newmd5value, err := LoadFileMD5(tempFile.Name())
	if err != nil {
		return 0, err
	}

	if newmd5value != md5value {
		return 0, fmt.Errorf("pls purge file %v origin md5 %v, with new_md5 %v", filename, md5value, newmd5value)
	}

	err = opdb.UpdateFileDetails(md5value, filename, tempFile.Name(), true)
	if err != nil {
		return 0, err
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

func MergeFile(c *gin.Context) {

}
