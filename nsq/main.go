package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

var (
	topicArray    = []string{"testTopic"} //, "drawpush", "optinfo", "ruokdrawprizetnum"}
	nsqlookupDURI = "127.0.0.1:4161"

	// 宣告一個共用的chan來做傳遞數據用
	msgChan chan []byte
)

func newNsqConsumer(channel string, topic string) *nsq.Consumer {
	config := nsq.NewConfig()

	q, err := nsq.NewConsumer(topic, channel, config)
	if err != nil {
		panic(err)
	}
	return q
}

func nsqConnectURI(uri string, port string) string {
	return uri + ":" + port
}

type customerStruct struct{}

func (*customerStruct) HandleMessage(message *nsq.Message) error {
	// log.Printf("Got a message with : %v\n", string(message.Body))
	msgChan <- message.Body
	return nil
}

func consumeMsg() {
	go consume()
}

func consume() {
	for {
		select {
		case msg := <-msgChan:
			go printmsgchan(msg)
		}
	}
}

func printmsgchan(msg []byte) {
	log.Println(string(msg))
}

func newCustomerHandler() nsq.Handler {
	return &customerStruct{}
}

func main() {
	// chan 必須要先make才會有記憶體lcoate
	msgChan = make(chan []byte, 100)
	for _, topicName := range topicArray {
		nsqC := newNsqConsumer(topicName, topicName)
		nsqC.AddHandler(newCustomerHandler())

		err := nsqC.ConnectToNSQLookupd(nsqlookupDURI)
		if err != nil {
			log.Panic("Could not connect")
		}
	}

	go consumeMsg()

	gracefulShutdown()
}

func gracefulShutdown() {

	// create one chan to print awaiting signal on console
	sigs := make(chan os.Signal, 1)
	// create another chan to receive signal to interrupt original chan
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
