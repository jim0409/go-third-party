package natsdk

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
	"github.com/nats-io/stan.go/pb"
)

type Publisher interface {
	SendMsg(string, []byte) error
}

type Subscriber interface {
	RecvMsg(string) (stan.Subscription, error)
}

type Stan interface {
	Closed()
	Publisher
	Subscriber
}

type Client struct {
	Nats stan.Conn
}

func InitNatsServer(cid, url, usr, pwd string) (Stan, error) {
	nc, err := nats.Connect(url, nats.UserInfo(usr, pwd))
	if err != nil {
		return nil, err
	}
	// defer nc.Close()
	// clusterID, clientID
	sc, err := stan.Connect("stan", cid, stan.NatsConn(nc))
	if err != nil {
		return nil, err
	}

	return &Client{
		Nats: sc,
	}, nil
}

func (c *Client) Closed() {
	c.Nats.Close()
}

func (c *Client) SendMsg(subj string, bs []byte) error {
	return c.Nats.Publish(subj, bs)
}

func (c *Client) RecvMsg(subj string) (stan.Subscription, error) {
	startOpt := stan.StartAt(pb.StartPosition_NewOnly)
	// if cur != 0 {
	// 	startOpt = stan.StartAtSequence(cur)
	// } else if deliverLast {
	// 	startOpt = stan.StartWithLastReceived()
	// } else if deliverAll {
	// 	log.Print("subscribing with DeliverAllAvailable")
	// 	startOpt = stan.DeliverAllAvailable()
	// } else if startDelta != "" {
	// 	ago, err := time.ParseDuration(startDelta)
	// 	if err != nil {
	// 		sc.Close()
	// 		log.Fatal(err)
	// 	}
	// 	startOpt = stan.StartAtTimeDelta(ago)
	// }
	i := 0
	qgroup := ""
	durable := ""
	mcb := func(msg *stan.Msg) {
		i++
		// printMsg(msg, i)
		handleReceiveMsg(msg, i)
	}
	// sub, err := c.Nats.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
	// if err != nil {
	// 	sub.Unsubscribe()
	// 	return nil, err
	// }

	// return sub, nil
	return c.Nats.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
}

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var x int32
	binary.Read(bytesBuffer, binary.LittleEndian, &x)
	return int(x)
}

func handleReceiveMsg(m *stan.Msg, i int) {
	fmt.Println("================", i)
	fmt.Printf("opcode: %v\n", m.Data[0:1])                    // 1 byte
	fmt.Printf("source code: %v\n", m.Data[1:2])               // 2 byte
	fmt.Printf("payload size: %d\n", BytesToInt(m.Data[3:10])) // 8 byte
	// fmt.Printf("%s\n", m.Data[2:2])
	// fmt.Printf("%s\n", m.Data[3:3])
	// fmt.Printf("%s\n", m.Data[11:])
	data := m.Data[11:]
	var jsonObj = make(map[string]interface{})
	err := json.Unmarshal(data, &jsonObj)
	if err != nil {
		panic(err)
	}

	for k, v := range jsonObj {
		fmt.Println(k, v)
	}
}

func printMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, m)
}
