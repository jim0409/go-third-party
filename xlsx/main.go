package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx/v3"
)

func main() {
	// open an existing file
	wb, err := xlsx.OpenFile("./samplefile.xlsx")
	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	fmt.Println("Sheets in this file:")
	for i, sh := range wb.Sheets {
		fmt.Println(i, sh.Name)
	}
	fmt.Println("----")

	router := gin.New()
	router.LoadHTMLGlob("templates/*")
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/", homeHandler)
	router.GET("/excel-download", serveExcel)

	router.Run(":3000")
}

func homeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", nil)
}

func serveExcel(c *gin.Context) {
	file := xlsx.NewFile()
	for i := 0; i < 10; i++ {
		sheet, err := file.AddSheet(fmt.Sprintf("sheet-%d", i))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		row := sheet.AddRow()
		cell := row.AddCell()
		cell.Value = "I am a cell!"
	}
	var b bytes.Buffer
	if err := file.Write(&b); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	downloadName := time.Now().UTC().Format("data-20060102150405.xlsx")
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}
