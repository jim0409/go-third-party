package utils

import (
	"bytes"
	"encoding/binary"
)

const (
	ConstHeader       = "testHeader"
	ConstHeaderLength = len(ConstHeader)
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
	prefixLen := append([]byte(ConstHeader), IntToBytes(len(message))...)
	// append message body
	return append(prefixLen, message...)
	// return append(append([]byte(ConstHeader), IntToBytes(len(message))...), message...)
}

// Depack would handle `each` packet as a block of buffer
func Depack(buffer []byte) []byte {
	bufLen := len(buffer)
	if bufLen == 0 {
		return make([]byte, 0)
	}

	prefLen := ConstHeaderLength + ConstMLength
	data := make([]byte, 32)
	for i := 0; i < bufLen; i++ {
		if bufLen < i+prefLen {
			break
		}

		if string(buffer[i:i+ConstHeaderLength]) != ConstHeader {
			continue
		}

		messageLength := BytesToInt(buffer[i+ConstHeaderLength : i+prefLen])
		if bufLen < i+prefLen+messageLength {
			break
		}

		data = buffer[i+prefLen : i+prefLen+messageLength]
	}

	return data
}
