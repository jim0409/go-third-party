package packet

import "context"

type Type byte

const (
	_            Type = iota
	Handshake         = 0x01 // 發起握手
	HandshakeAck      = 0x02 // 發起握手ack
	HeartBeat         = 0x03 // 心跳包
	Data              = 0x04 // 資料包
	Kick              = 0x05 // 伺服器踢出玩家

	HeadLength    = 4
	MaxPacketSize = 64 * 1024
)

type Packet struct {
	Type    Type
	Length  int
	Data    []byte
	Context context.Context
}
