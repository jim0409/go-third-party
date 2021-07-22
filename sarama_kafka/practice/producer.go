package main

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type KafkaProducerImpl interface {
	PushToTopic(string, []byte, []byte)
	Close() error
}

type KafkaConfig struct {
	Brokers  []string
	RetryMax int
}

type KafkaClient struct {
	Producer sarama.AsyncProducer
}

func (c *KafkaClient) Close() error {
	return c.Producer.Close()
}

func (c *KafkaClient) PushToTopic(topic string, keydata []byte, valuedata []byte) {
	// 如果要用單一有序發送，用 NewSyncProducer；平行處理用 NewAsyncProducer
	c.Producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(keydata),
		Value: sarama.StringEncoder(valuedata),
	}

	// 接收到成功的消息
	msg := <-c.Producer.Successes()
	fmt.Println(msg, "--- offset --- ", msg.Offset)
}

func NewKafakaConfig(brokers []string, retryTime int) *KafkaConfig {
	return &KafkaConfig{
		Brokers:  brokers,
		RetryMax: retryTime,
	}
}

func NewKafkaProducerImpl(c *KafkaConfig) (KafkaProducerImpl, error) {
	config := sarama.NewConfig()
	config.Producer.Retry.Max = c.RetryMax
	// config.Producer.RequiredAcks = sarama.NoResponse
	config.Producer.RequiredAcks = sarama.WaitForAll
	// config.Producer.RequiredAcks = sarama.WaitForLocal
	/*
		不設定 true 會導致 Producer.Successes() block .. 且無法保證 Producer 不掉落訊息;
		https: //pkg.go.dev/github.com/Shopify/sarama#example-AsyncProducer-Goroutines
	*/
	config.Producer.Return.Successes = true

	producer, err := sarama.NewAsyncProducer([]string(c.Brokers), config)
	if err != nil {
		return nil, fmt.Errorf("Error occur while create Kafka Producer :%v", err)
	}

	return &KafkaClient{Producer: producer}, nil
}

func main() {

	topic := "sarama"

	// broker1 := "127.0.0.1:9091"
	broker2 := "127.0.0.1:9092"
	// broker3 := "127.0.0.1:9093"
	brokers := []string{broker2}
	retryMax := 5

	kc := NewKafakaConfig(brokers, retryMax)
	kpI, err := NewKafkaProducerImpl(kc)
	if err != nil {
		panic(err)
	}

	key := "123"
	value := "456"
	kpI.PushToTopic(topic, []byte(key), []byte(value))

	kpI.Close()
}
