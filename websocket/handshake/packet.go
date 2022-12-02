package main

import (
	"context"
)

type Type byte

const (
	_            Type = iota
	Handshake         = 0x01 // 發起握手
	HandshakeAck      = 0x02 // 發起握手ack
	HeartBeat         = 0x03 // 心跳包
	Data              = 0x04 // 資料包
	Kick              = 0x05 // 伺服器踢出玩家
	TraceData         = 0x06 // 帶有 trace 的 data
)

type Packet struct {
	Typ     Type
	Length  int
	Data    []byte
	Context context.Context
}

// Encode packet data length to bytes(Big end)
func intToBytes(n int) []byte {
	buf := make([]byte, 3)
	buf[0] = byte((n >> 16) & 0xFF)
	buf[1] = byte((n >> 8) & 0xFF)
	buf[2] = byte(n & 0xFF)
	return buf
}

// Decode packet data length byte to int(Big end)
func bytesToInt(b []byte) int {
	result := 0
	for _, v := range b {
		result = result<<8 + int(v)
	}
	return result
}

// Protocol Frame
// -<type>-|-------<lenght>-------|-<data>-
// 1 byte packet type,
// 3 bytes packet data length(big end), and data segment
func Encode(typ *Packet) {}
