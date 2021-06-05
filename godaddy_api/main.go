package main

import (
	"flag"
	"fmt"
	"os"
)

const godaddyProdUrl = "https://api.godaddy.com/v1"

var (
	csvFilePath  = flag.String("file", "godaddy_ac.csv", "specific the input csv file path")
	checkcommit  = flag.Bool("version", false, "burry code for check version")
	gitcommitnum string
)

func checkComimit() {
	fmt.Println(gitcommitnum)
}

func main() {
	flag.Parse()
	if *checkcommit {
		checkComimit()
		os.Exit(1)
	}

	userInfos := LoadAccountFromCSV(*csvFilePath)
	if err := userInfos.RetriveAccountDomain(); err != nil {
		panic(err)
	}

	if err := userInfos.DumpDataToCsv(); err != nil {
		panic(err)
	}

	fmt.Println("done")
}
