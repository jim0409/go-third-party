package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

type UserInfo struct {
	Id         string
	Name       string
	CustomerId string
	Key        string
	Sec        string
	Domains    []Domain
}

type Domain struct {
	Name       string
	Expiretime string
}

type UserInfos struct {
	UserInfos []UserInfo
}

type CustomerBehavior interface {
	RetriveAccountDomain() error
	RetriveExpireDomain() error
	DumpDataToCsv() error
}

func LoadAccountFromCSV(csvFile string) CustomerBehavior {
	return &UserInfos{
		UserInfos: readCsvFile(csvFile),
	}
}

func (u *UserInfos) RetriveAccountDomain() error {
	// assume the first row is index only
	for i := 1; i < len(u.UserInfos); i++ {
		// fmt.Printf("%v___%v\n", i, u.UserInfos[i])
		m, err := domains(u.UserInfos[i].CustomerId, u.UserInfos[i].Key, u.UserInfos[i].Sec)
		if err != nil {
			return err
		}

		for index, keyValue := range m {
			fmt.Println(index)

			tmpDomain := Domain{
				Name:       fmt.Sprintf("%v", keyValue["domain"]),
				Expiretime: fmt.Sprintf("%v", keyValue["expires"]),
			}
			// fmt.Println(index, keyValue)

			u.UserInfos[i].Domains = append(u.UserInfos[i].Domains, tmpDomain)
		}
	}

	return nil
}

func (u *UserInfos) RetriveExpireDomain() error {
	return nil
}

func (u *UserInfos) DumpDataToCsv() error {
	data := [][]string{{"ID", "Name", "CustomerId", "DomainName", "ExpireDate"}}
	for i, j := range u.UserInfos {
		_ = i
		// fmt.Println(i, j)
		for id, dd := range j.Domains {
			tmpStrArr := []string{j.Id, j.Name, j.CustomerId, dd.Name, dd.Expiretime}
			_ = id
			// fmt.Println(id, dd)
			// data[1][1] = append(data[1][1], tmpStrArr)
			data = append(data, tmpStrArr)
		}
	}

	return write("result.csv", data)
}

// refer: https://medium.com/@ankurraina/reading-a-simple-csv-in-go-36d7a269cecd
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

// refer: https://golangcode.com/write-data-to-a-csv-file/
func write(filename string, data [][]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Cannot create file: %v", err)

	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		if err != nil {
			return fmt.Errorf("Cant write to file %v", err)
		}
	}

	return err
}
