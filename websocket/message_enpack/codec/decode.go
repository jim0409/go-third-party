package codec

import (
	"bytes"
	"context"
	"errors"
	"go-third-party/websocket/message_enpack/packet"
)

const (
	HeadLength    = 4
	MaxPacketSize = 64 * 1024
)

// ErrPacketSizeExcced is the error used for encode/decode.
var ErrPacketSizeExcced = errors.New("codec: packet size exceed")

// A Decoder reads and decodes network data slice
type Decoder struct {
	buf  *bytes.Buffer
	size int  // last packet length
	typ  byte // last packet type
}

// NewDecoder returns a new decoder that used for decode network bytes slice.
func NewDecoder() *Decoder {
	return &Decoder{
		buf:  bytes.NewBuffer(nil),
		size: -1,
	}
}

// Decode decode the network bytes slice to packet.Packet(s)
// TODO(Warning): shared slice
func (c *Decoder) Decode(data []byte) ([]*packet.Packet, error) {
	c.buf.Write(data)

	var (
		packets []*packet.Packet
		err     error
	)
	// check length
	if c.buf.Len() < HeadLength {
		return nil, err
	}

	// first time
	if c.size < 0 {
		if err = c.forward(); err != nil {
			return nil, err
		}
	}

	for c.size <= c.buf.Len() {
		p := &packet.Packet{
			Type:    packet.Type(c.typ),
			Length:  c.size,
			Data:    c.buf.Next(c.size),
			Context: context.Background(),
		}
		packets = append(packets, p)

		// more packet
		if c.buf.Len() < HeadLength {
			c.size = -1
			break
		}

		if err = c.forward(); err != nil {
			return nil, err

		}

	}

	return packets, nil
}

func (c *Decoder) forward() error {
	header := c.buf.Next(HeadLength)
	c.typ = header[0]
	if c.typ < packet.Handshake || c.typ > packet.Kick {
		return ErrWrongPacketType
	}
	c.size = bytesToInt(header[1:])

	// packet length limitation
	if c.size > MaxPacketSize {
		return ErrPacketSizeExcced
	}
	return nil
}

// Decode packet data length byte to int(Big end)
func bytesToInt(b []byte) int {
	result := 0
	for _, v := range b {
		result = result<<8 + int(v)
	}
	return result
}
