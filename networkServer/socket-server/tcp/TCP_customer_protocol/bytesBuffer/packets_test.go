package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"testing"
)

func TestEnpack(t *testing.T) {
	bytesBuffer := &bytes.Buffer{}
	array := make([]byte, 20)
	array[0] = byte(1)
	// array[1] = byte(2 >> 8)
	array[1] = byte(2)
	array[2] = byte(33332 >> 8)
	// bytesBuffer.WriteString("a" >> 8)
	// bytesBuffer.WriteRune(array)
	for _, j := range array {
		bytesBuffer.WriteByte(j)
	}

	log.Printf("%v\n", bytesBuffer.Bytes())

}

func TestIntToBytes(t *testing.T) {
	intArr := []int{1, 2, 3, 4, 100}
	for _, j := range intArr {
		x := int32(j)
		bytesBuffer := bytes.NewBuffer([]byte{})
		// bytesBuffer := &bytes.Buffer{}
		err := binary.Write(bytesBuffer, binary.BigEndian, x)
		if err != nil {
			log.Fatal(err)
		}
		// log.Println(bytesBuffer)
		// bytesBuffer.WriteRune(j)
		log.Println(bytesBuffer.Bytes())
	}
}
