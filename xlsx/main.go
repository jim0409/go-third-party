package main

import (
	"fmt"

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

}
