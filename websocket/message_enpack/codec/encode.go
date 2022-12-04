package codec

import (
	"context"
	"errors"
	"go-third-party/websocket/message_enpack/packet"
)

var ErrWrongPacketType = errors.New("wrong packet type")

// Protocol Frame
// -<type>-|-------<lenght>-------|-<data>-
// 1 byte packet type,
// 3 bytes packet data length(big end), and data segment
func Encode(typ packet.Type, data []byte) ([]byte, error) {
	if typ < Handshake || typ > Kick {
		return nil, ErrWrongPacketType
	}

	p := &packet.Packet{
		Type:    typ,
		Length:  len(data),
		Context: context.Background(),
	}
	buf := make([]byte, p.Length+HeadLength)
	buf[0] = byte(p.Type)

	copy(buf[1:HeadLength], intToBytes(p.Length))
	copy(buf[HeadLength:], data)

	return buf, nil
}

// Encode packet data length to bytes(Big end)
func intToBytes(n int) []byte {
	buf := make([]byte, 3)
	buf[0] = byte((n >> 16) & 0xFF)
	buf[1] = byte((n >> 8) & 0xFF)
	buf[2] = byte(n & 0xFF)
	return buf
}
