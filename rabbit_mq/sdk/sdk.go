package sdk

import (
	"fmt"

	"github.com/streadway/amqp"
)

type Client struct {
	Channel *amqp.Channel
}

type IClient interface {
}

func NewAMQP(usr, pwd, addr, port string) (IClient, error) {
	dst := fmt.Sprintf("amqp://%v:%v@%v:%v", usr, pwd, addr, port)
	conn, err := amqp.Dial(dst)
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}

	return &Client{
		Channel: ch,
	}, nil
}

func (c *Client) Consume(name string) error {

	msgs, err := c.Channel.Consume(
		name,  // queue
		"",    // consumer
		true,  // auto-ack
		false, // exclusive
		false, // no-local
		false, // no-wait
		nil,   // args
	)

	if err != nil {
		return err
	}

	go func() {
		for d := range msgs {
			fmt.Printf("Received msg: %s", d.Body)
		}
	}()

	forever := make(chan bool)
	fmt.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

	return nil
}
