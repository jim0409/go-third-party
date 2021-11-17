package main

import (
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
		printMsg(msg, i)
	}
	// sub, err := c.Nats.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
	// if err != nil {
	// 	sub.Unsubscribe()
	// 	return nil, err
	// }

	// return sub, nil
	return c.Nats.QueueSubscribe(subj, qgroup, mcb, startOpt, stan.DurableName(durable))
}

func printMsg(m *stan.Msg, i int) {
	log.Printf("[#%d] Received on [%s]: '%s'\n", i, m.Subject, m)
}
