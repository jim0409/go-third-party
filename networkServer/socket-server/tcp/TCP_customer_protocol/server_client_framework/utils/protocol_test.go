package utils

import (
	"fmt"
	"testing"
)

func TestIntToBytes(t *testing.T) {
	a := 1
	ba := IntToBytes(a)

	// fmt.Println(a)
	fmt.Printf("%v\n", ba)
}
