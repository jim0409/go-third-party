package utils

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader       = "testHeader"
	ConstHeaderLength = 10
	ConstMLength      = 4
)

// IntToBytes 轉換int為[]byte: 透過binary包，Write方法，將int寫為bytesBuffer
/*
	transfer int(n) as int32
	declare x as int32(n)
	Write x as bytesBuffer(:=bytes.NewBuffer([]byte{}))
*/
func IntToBytes(n int) []byte {
	x := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, x)
	return bytesBuffer.Bytes()
}

// BytesToInt 轉換[]byte為int: 透過binary包，Read方法，將bytesBuffer轉換為int
/*
	declare bytesBuffer as bytes.NewBuffer(b)
	map bytesBuffer to int32
*/
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.BigEndian, &x)
	return int(x)
}

func Enpack(message []byte) []byte {
	return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

func Depack(buffer []byte) []byte {
	length := len(buffer)

	var i int
	data := make([]byte, 32)
	for i = 0; i < length; i = i + 1 {
		if length < i+ConstHeaderLength+ConstMLength {
			break
		}
		if string(buffer[i:i+ConstHeaderLength]) == ConstHeader {
			messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+ConstHeaderLength+ConstMLength])
			if length < i+ConstHeaderLength+ConstMLength+messageLength {
				break
			}
			data = buffer[i+ConstHeaderLength+ConstMLength : i+ConstHeaderLength+ConstMLength+messageLength]

		}
	}

	if i == length {
		return make([]byte, 0)
	}
	return data
}
