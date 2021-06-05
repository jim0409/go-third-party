package main

import (
	"flag"
	"log"
	"sync"

	"github.com/nsqio/go-nsq"
)

var (
	chanFlag           = flag.String("channel", "ch", "pre-define the consumer channel.")
	topicFlag          = flag.String("topic", "write_test", "pre-define the consumer topic.")
	nsqAddr            = flag.String("nsqAddr", "127.0.0.1", "Specific the nsq address to write")
	nsqPort            = flag.String("nsqPort", "4150", "Specify the used nsqd port")
	nsqPollingInterval = flag.Int("nsqT", 10, "Specify the polling interval for consumer __default: 10ms")

	// main controll variable for nsq
	nsqC   = &nsq.Consumer{}
	nsqURI string
	wg     = &sync.WaitGroup{}
)

func newNsqConsumer(topic string, channel string) *nsq.Consumer {
	config := nsq.NewConfig()
	// lookupd_poll_interval" min:"10ms" max:"5m" default:"60s
	// config.LookupdPollInterval = time.Duration(int(*nsqPollingInterval)) * time.Millisecond

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
	log.Printf("Got a message with : %v\n", string(message.Body))

	// wg.Done() // add this func to break the processs after recev a msg from nsq queue
	return nil
}

func newCustomerHandler() nsq.Handler {
	return &customerStruct{}
}

func main() {

	// wg := &sync.WaitGroup{}  // let wg be global variable to controll the customer handler.
	wg.Add(1)

	// nsqC.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
	// 	log.Printf("Got a message: %v", message)
	// 	wg.Done()
	// 	return nil
	// }))
	nsqC.AddHandler(newCustomerHandler())

	err := nsqC.ConnectToNSQD(nsqURI)
	if err != nil {
		log.Panic("Could not connect")
	}
	wg.Wait()

}

func init() {
	flag.Parse()
	nsqC = newNsqConsumer(*topicFlag, *chanFlag)
	nsqURI = nsqConnectURI(*nsqAddr, *nsqPort)
}
