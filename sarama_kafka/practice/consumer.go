package main

import (
	"context"
	"log"
	"os"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	Topics        []string
	ConsumerGroup sarama.ConsumerGroup
}

// 在運行 consumer 以前會使用 setup
func (c *Consumer) Setup(s sarama.ConsumerGroupSession) error {
	// s.MarkOffset(c.Topics)
	return nil
}

// 當 consumer 運行結束後執行 cleanup
func (c *Consumer) Cleanup(s sarama.ConsumerGroupSession) error {
	// s.Commit()
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: key = %s, value = %v, topic = %s, partition = %v, offset = %v", string(message.Key), string(message.Value), message.Topic, message.Partition, message.Offset)
		session.MarkMessage(message, "")
		session.Commit() // 要是沒有 commit 就會對於該 consumer gorup 重複消費!
	}

	return nil
}

func (c *Consumer) Execute(topics []string) chan error {
	var echan = make(chan error, 1)
	go func() {
		for {
			if err := c.ConsumerGroup.Consume(context.TODO(), topics, c); err != nil {
				// log.Panicf("Error from consumer: %v", err)
				echan <- err
				return
			}

		}
	}()

	return echan
}

type ConsumerImp interface {
	// must satisfy the ConsumerGroupHandle interface {Setup, Cleanup, ConsumeClaim}
	Setup(sarama.ConsumerGroupSession) error
	Cleanup(sarama.ConsumerGroupSession) error
	ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error

	// start loop consume msg via consumer group
	Execute([]string) chan error
}

func NewConsumeHandler(topics []string, brokers []string, group string) ConsumerImp {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true
	config.Consumer.Offsets.AutoCommit.Enable = false
	client, err := sarama.NewConsumerGroup(brokers, group, config)
	if err != nil {
		log.Panicf("Error creating consumer group client: %v", err)
	}

	return &Consumer{
		Topics:        topics,
		ConsumerGroup: client,
	}
}

func main() {
	sarama.Logger = log.New(os.Stdout, "[sarama - practice]", log.LstdFlags)
	topics := []string{"sarama"}
	broker := "127.0.0.1:9092"
	group := "test"

	// declare the consumer handler
	consumer := NewConsumeHandler(topics, []string{broker}, group)
	errs := consumer.Execute(topics)
	for err := range errs {
		log.Fatal(err)
	}
}
