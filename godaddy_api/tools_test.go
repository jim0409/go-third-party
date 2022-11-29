package main

import (
	"testing"
)

func TestReadCsv(t *testing.T) {
	file := "csv_for_test.csv"
	us := readCsvFile(file)

	// count from the `1st` ac ... since the `0` initial row is just for name
	u := (us)[1]
	assert.Equal(t, "1", u.Id)
	assert.Equal(t, "jim", u.Name)
	assert.Equal(t, "280076842", u.CustomerId)
	assert.Equal(t, "e5N8Yzzn7hBo_Fxe7rVEZwvyTfMW4ztQmJG", u.Key)
	assert.Equal(t, "PDnDKUFDZjefBUUyTeWGkb", u.Sec)
}

func TestWriteCsv(t *testing.T) {
	var data = [][]string{{"Line1", "Hello Readers of"}, {"Line2", "golangcode.com"}}
	err := write("result.csv", data)
	assert.NilError(t, err)
}

func TestLoadAccountFromCSV(t *testing.T) {
	uinfos := LoadAccountFromCSV("csv_for_test.csv")
	uinfos.RetriveAccountDomain()
	uinfos.DumpDataToCsv()
}
