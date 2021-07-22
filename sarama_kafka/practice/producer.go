package main

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/Shopify/sarama"
)

type KafkaProducerImpl interface {
	PushToTopic(string, []byte, []byte) chan interface{}
	// PushToTopic(string, []byte, []byte) chan *sarama.ProducerMessage
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

func (c *KafkaClient) PushToTopic(topic string, keydata []byte, valuedata []byte) chan interface{} {
	// func (c *KafkaClient) PushToTopic(topic string, keydata []byte, valuedata []byte) chan *sarama.ProducerMessage {
	out := make(chan interface{}, 1)
	// out := make(chan *sarama.ProducerMessage, 1)
	// 如果要用單一有序發送，用 NewSyncProducer；平行處理用 NewAsyncProducer
	c.Producer.Input() <- &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(keydata),
		Value: sarama.StringEncoder(valuedata),
	}

	// 接收到成功的消息
	// msg := <-c.Producer.Successes()
	// fmt.Println(msg, "--- offset --- ", msg.Offset)
	// msg := c.Producer.Successes()
	out <- c.Producer.Successes()
	// fmt.Println(reflect.TypeOf(c.Producer.Successes()))
	// out <- msg
	return out

	// return out
}

func NewKafakaConfig(brokers []string, retryTime int) *KafkaConfig {
	return &KafkaConfig{
		Brokers:  brokers,
		RetryMax: retryTime,
	}
}

func NewKafkaProducerImpl(c *KafkaConfig) (KafkaProducerImpl, error) {
	config := sarama.NewConfig()
	// config.Producer.Partitioner = sarama.NewManualPartitioner
	// config.Producer.Partitioner = sarama.NewRandomPartitioner
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
			// fmt.Println(<-o)
			wg.Done()
		}()
	}

	// var i = 1
	go func() {
		for msg := range out {
			// fmt.Println(reflect.TypeOf(msg))
			if v, ok := msg.(chan interface{}); ok {
				// fmt.Println(reflect.TypeOf(v))
				// fmt.Println(v)
				v1 := <-v
				// fmt.Println(reflect.TypeOf(v1))
				if vv, ok := v1.(<-chan *sarama.ProducerMessage); ok {
					vvv := <-vv
					fmt.Println(reflect.TypeOf(vvv))
					fmt.Println(vvv, "========", vvv.Offset)
					// fmt.Println(vv)
				} else {
					fmt.Println("XXX")
				}
			} else {
				fmt.Println("???")
			}

		}
	}()

	wg.Wait()

	kpI.Close()
}
