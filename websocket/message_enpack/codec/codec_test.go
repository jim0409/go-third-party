package codec

import (
	"encoding/hex"
	"fmt"
	"testing"

	"go-third-party/websocket/message_enpack/packet"

	"github.com/stretchr/testify/assert"
)

func TestEncode(t *testing.T) {
	str := []byte("hello")
	bs, err := Encode(packet.Data, str)
	assert.Nil(t, err)

	fmt.Println(hex.EncodeToString(str))
	fmt.Println(hex.EncodeToString(bs))

}

func TestDecode(t *testing.T) {
	str := []byte("hello")
	bs, err := Encode(packet.Data, str)
	assert.Nil(t, err)

	decoder := NewDecoder()
	pkts, err := decoder.Decode(bs)
	assert.Nil(t, err)
	for _, pkt := range pkts {
		fmt.Println(pkt.Type)
		fmt.Println(pkt.Length)
		fmt.Println(string(pkt.Data))
		fmt.Println(pkt.Context)
	}
}
