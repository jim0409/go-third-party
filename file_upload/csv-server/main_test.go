package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"testing"
)

type UserInfo struct {
	Id         string
	Name       string
	CustomerId string
	Key        string
	Sec        string
}

// test for reading csv file
func readCsvFile(f string) []UserInfo {
	csvfile, err := os.Open(f) // Open the file
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	// Parse the file
	r := csv.NewReader(csvfile) // or ... r := csv.NewReader(bufio.NewReader(csvfile))

	UserInfos := make([]UserInfo, 0) // Read each record from csv
	for {
		record, err := r.Read() // Read each record from csv
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		u := UserInfo{
			Id:         record[0],
			Name:       record[1],
			CustomerId: record[2],
			Key:        record[3],
			Sec:        record[4],
		}
		UserInfos = append(UserInfos, u)
	}
	return UserInfos
}

func TestUpload(t *testing.T) {
	readCsvFile("files/upload.csv")
}
