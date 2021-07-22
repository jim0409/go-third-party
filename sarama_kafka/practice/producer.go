package main

import (
	"fmt"
	"sync"

	"github.com/Shopify/sarama"
)

type KafkaProducerImpl interface {
	PushToTopic(string, []byte, []byte) <-chan *sarama.ProducerMessage
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

func (c *KafkaClient) PushToTopic(topic string, keydata []byte, valuedata []byte) <-chan *sarama.ProducerMessage {
	c.Producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(keydata),
		Value: sarama.StringEncoder(valuedata),
	}

	// 接收到成功的消息
	return c.Producer.Successes()
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

	var wg sync.WaitGroup
	out := make(chan interface{})
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			out <- kpI.PushToTopic(topic, []byte(key), []byte(value))
			wg.Done()
		}()
	}

	go func() {
		for o := range out {
			if v, ok := o.(<-chan *sarama.ProducerMessage); ok {
				msg := <-v
				fmt.Println(msg)
			} else {
				fmt.Println("???")
			}
		}
	}()

	wg.Wait()

	kpI.Close()
}
