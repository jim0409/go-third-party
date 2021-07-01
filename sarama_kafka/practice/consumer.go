package main

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

type Consumer struct {
	Topics        []string
	ConsumerGroup sarama.ConsumerGroup
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		log.Printf("Message claimed: key = %s, value = %v, topic = %s, partition = %v, offset = %v", string(message.Key), string(message.Value), message.Topic, message.Partition, message.Offset)
		session.MarkMessage(message, "")
	}

	return nil
}

func (c *Consumer) Execute(topics []string) error {
	// var err error
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

	return <-echan
}

type ConsumerImp interface {
	// must satisfy the ConsumerGroupHandle interface {Setup, Cleanup, ConsumeClaim}
	Setup(sarama.ConsumerGroupSession) error
	Cleanup(sarama.ConsumerGroupSession) error
	ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error

	// start loop consume msg via consumer group
	Execute([]string) error
}

func NewConsumeHandler(topics []string, brokers []string, group string) ConsumerImp {
	config := sarama.NewConfig()
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
	topics := []string{"sarama"}
	broker := "127.0.0.1:9092"
	group := "test"

	// declare the consumer handler
	consumer := NewConsumeHandler(topics, []string{broker}, group)
	if err := consumer.Execute(topics); err != nil {
		log.Fatal(err)
	}
}
