package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context) {
	// TODO: retrieve user from JWT
	// verified upload permission

	// md5value
	md5value := c.Query("md5value")
	if md5value == "" {
		c.JSON(400, "lack of md5value!")
		return
	}

	// filename
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(400, "lack of filename!")
		return
	}

	// totalchunks
	stotalchunks := c.Query("totalchunks")
	if stotalchunks == "" {
		c.JSON(400, "lack of totalchunks!")
		return
	}

	totalchunks, err := strconv.Atoi(stotalchunks)
	if err != nil {
		c.JSON(400, fmt.Sprintf("convert totalchunks err %v!", err))
		return
	}

	// chunkorder
	schunkorder := c.Query("chunkorder")
	if schunkorder == "" {
		c.JSON(400, "lack of chunkorder!")
		return
	}
	chunkorder, err := strconv.Atoi(schunkorder)
	if err != nil {
		c.JSON(400, fmt.Sprintf("convert chunkorder err %v!", err))
		return
	}

	// username
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

	// 如果已經上傳則直接回傳
	if uploadfile.ID != 0 && uploadfile.IsUploaded == 1 {
		switch uploadfile.IsUploaded {
		case 1:
			c.JSON(200, gin.H{
				"status":   "success",
				"id":       md5value,
				"filename": fmt.Sprintf("Uploaded File: %+v", uploadfile.FileName),
				"size":     fmt.Sprintf("File Size: %+v", uploadfile.Size),
			})
		case 2:
			c.JSON(200, gin.H{
				"status":   "uploading",
				"id":       md5value,
				"filename": fmt.Sprintf("Uploaded File: %+v", uploadfile.FileName),
				"size":     fmt.Sprintf("File Size: %+v", uploadfile.Size),
			})

		// 考慮用排程系統處理刪除失敗的檔案, 避免過多的刪除失敗導致穿透
		case -1:
			c.JSON(200, gin.H{
				"status":   "failed",
				"id":       md5value,
				"filename": fmt.Sprintf("Uploaded File: %+v", uploadfile.FileName),
				"size":     fmt.Sprintf("File Size: %+v", uploadfile.Size),
			})
		}
		return
	}

	r := c.Request
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		fmt.Printf("Error Retrieving the File %v\n", err)
		return
	}
	defer file.Close()

	// TODO: if size over restrict, return upload error msg
	size := handler.Size

	//	尚未上傳: uploadfile.ID 為空 or uploadfile.IsUploaded = 0
	id, err := BackUpFile(file, username, filename, md5value, size, chunkorder, totalchunks)
	if err != nil {
		c.JSON(404, fmt.Sprintf("Failed to Uploaded File %v", err))
		return
	}

	c.JSON(200, gin.H{
		"status":   "success",
		"id":       id,
		"filename": fmt.Sprintf("Uploaded File: %+v", handler.Filename),
		"size":     fmt.Sprintf("File Size: %+v", handler.Size),
	})
}

func BackUpFile(file io.Reader, username string, filename string, md5value string, size int64, chunknum int, totalchunk int) (int, error) {
	id, err := opdb.InsertOneRecord(username, filename, md5value, size, totalchunk)
	if err != nil {
		return 0, err
	}

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return 0, err
	}

	chunkFileName := fmt.Sprintf("files/%s_%s_%d_%d", username, filename, totalchunk, chunknum)

	f, err := os.Create(chunkFileName)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	_, err = f.Write(fileBytes)
	if err != nil {
		return 0, err
	}

	// can't use tempFile since in memory bytes differ from file md5
	newmd5value, err := LoadFileMD5(chunkFileName)
	if err != nil {
		return 0, err
	}

	if newmd5value != md5value {
		return 0, fmt.Errorf("pls purge file %v origin md5 %v, with new_md5 %v", filename, md5value, newmd5value)
	}

	err = opdb.UpdateFileDetails(md5value, filename, chunkFileName, chunknum)
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

type ChunkFileMd5Value struct {
	Md5Values []string `json:"chunk_file_md5"`
}

func MergeFile(c *gin.Context) {
	// filename
	filename := c.Query("filename")
	if filename == "" {
		c.JSON(400, "lack of filename!")
		return
	}

	// username
	username := c.GetHeader("username")
	if username == "" {
		c.JSON(400, "lack of username!")
		return
	}

	chunkmd5s := ChunkFileMd5Value{}
	bs, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(400, err)
		return
	}

	err = json.Unmarshal(bs, &chunkmd5s)
	if err != nil {
		c.JSON(400, err)
		return
	}

	chunkfiles, err := opdb.FindUploadDetailByMd5Values(chunkmd5s.Md5Values)
	if err != nil {
		c.JSON(400, err)
		return
	}

	err = MergeChunkFiles(filename, username, chunkfiles)
	if err != nil {
		c.JSON(400, err)
		return
	}

}

func MergeChunkFiles(filename string, username string, chunkfiles *[]FileUploadDetail) error {
	// 儲存到本地
	saveFile := fmt.Sprintf("files/%v", filename)
	f, err := os.Create(saveFile)
	if err != nil {
		return err
	}
	defer f.Close()

	var size int64
	for _, chunkfile := range *chunkfiles {
		chunkbytes, err := os.ReadFile(chunkfile.ChunkFilename)
		if err != nil {
			return err
		}

		_, err = f.Write(chunkbytes)
		if err != nil {
			return err
		}

		size = size + chunkfile.Size
	}

	// 合併成功則在 FileList 增加一筆紀錄
	return opdb.AddFileToList(filename, username, size)
}
