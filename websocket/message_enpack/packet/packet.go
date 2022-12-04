package packet

import "context"

type Type byte

type Packet struct {
	Type    Type
	Length  int
	Data    []byte
	Context context.Context
}
