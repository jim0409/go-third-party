package main

import (
	"flag"
	"log"

	"github.com/nsqio/go-nsq"
)

var (
	chanFlag  = flag.String("channel", "ch", "pre-define the consumer channel.")
	topicFlag = flag.String("topic", "write_test", "pre-define the consumer topic.")
	nsqAddr   = flag.String("nsqAddr", "127.0.0.1", "Specific the nsq address to write")
	nsqPort   = flag.String("nsqPort", "4150", "Specify the used nsqd port")
	nsqMsg    = flag.String("nsqMsg", "test", "Enter the msg want to send to nsq")

	// main variables
	nsqP   = &nsqProducer{}
	nsqURI string
)

type nsqProducer struct {
	producer *nsq.Producer
	topic    string
	msg      string
}

func (nq *nsqProducer) Stop() {
	nq.producer.Stop()
}

func (nq *nsqProducer) Publish() error {
	var err error
	err = nq.producer.Publish(nq.topic, []byte(nq.msg))
	return err
}

func nsqConnectURI(uri string, port string) string {
	return uri + ":" + port
}

func newNsqProducer(channel string, topic string) *nsq.Producer {
	config := nsq.NewConfig()
	w, err := nsq.NewProducer(nsqURI, config)
	if err != nil {
		panic(err)
	}
	return w
}

func main() {
	// config := nsq.NewConfig()
	// w, _ := nsq.NewProducer("127.0.0.1:4150", config)

	// for i := 0; i < 100; i++ {

	err := nsqP.Publish()
	// err := nsqP.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}
	// }

	nsqP.Stop()
}

func init() {
	flag.Parse()
	nsqURI = nsqConnectURI(*nsqAddr, *nsqPort)
	nsqP.producer = newNsqProducer(*chanFlag, *topicFlag)
	nsqP.topic = *topicFlag
	nsqP.msg = *nsqMsg
}
