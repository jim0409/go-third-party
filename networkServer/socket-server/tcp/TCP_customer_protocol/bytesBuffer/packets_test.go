package main

import (
	"bytes"
	"encoding/binary"
	"log"
	"testing"
)

func TestEnpack(t *testing.T) {
	bytesBuffer := &bytes.Buffer{}
	// for each byte is uint8
	// ,meanwhile its only 256 for each value!
	array := make([]byte, 20)
	array[0] = byte(1)
	// array[1] = byte(2 >> 8)
	array[1] = byte(2)
	array[2] = byte(33332 >> 8) // divide 2^8
	array[3] = byte(10 << 2)    // multiple 2^2
	array[4] = byte(255)        // max number of each element
	array[5] = byte(0xff)       // replace number with 16 digits
	array[6] = byte(1 & 0xff)   // and logic with 0xff
	array[7] = byte(8 & 0xff)   // and with numeric number
	array[8] = byte(16 & 0xff)  //  ...
	array[9] = byte(124 | 0xff) // or with 0xff
	array[10] = byte(1 ^ 0xff)  // orNot with 0xff
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
