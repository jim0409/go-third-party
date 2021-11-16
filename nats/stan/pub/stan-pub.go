// Copyright 2016-2018 The NATS Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/stan.go"
)

var usageStr = `
Usage: stan-pub [options] <subject> <message>

Options:
	-s, --server   <url>            NATS Streaming server URL(s)
	-c, --cluster  <cluster name>   NATS Streaming cluster name
	-id,--clientid <client ID>      NATS Streaming client ID
	-a, --async                     Asynchronous publish mode
`

// NOTE: Use tls scheme for TLS, e.g. stan-pub -s tls://demo.nats.io:4443 foo hello
func usage() {
	fmt.Printf("%s\n", usageStr)
	os.Exit(0)
}

func main() {
	var clusterID string
	var clientID string
	var async bool
	var URL string
	var Usr string
	var Passwd string
	var subj string
	var message string

	flag.StringVar(&URL, "s", stan.DefaultNatsURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&URL, "server", stan.DefaultNatsURL, "The nats server URLs (separated by comma)")
	flag.StringVar(&clusterID, "c", "stan", "The NATS Streaming cluster ID")
	flag.StringVar(&clusterID, "cluster", "stan", "The NATS Streaming cluster ID")
	flag.StringVar(&clientID, "id", "stan-pub", "The NATS Streaming client ID to connect with")
	flag.StringVar(&clientID, "clientid", "stan-pub", "The NATS Streaming client ID to connect with")
	flag.BoolVar(&async, "a", false, "Publish asynchronously")
	flag.BoolVar(&async, "async", false, "Publish asynchronously")
	flag.StringVar(&Usr, "usr", "stan", "specific the nats streaming cluster coonect usr.")
	flag.StringVar(&Passwd, "pwd", "stan", "specific the nats streaming cluster connect usr password.")
	flag.StringVar(&subj, "sub", "stan", "specific the nats streaming cluster publish subject.")
	flag.StringVar(&message, "msg", "test", "sending msg ...")

	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	// args := flag.Args()

	// if len(args) < 1 {
	// 	usage()
	// }

	// URL = fmt.Sprintf("jim:password@%v", URL)
	nc, _ := nats.Connect(URL, nats.UserInfo(Usr, Passwd))
	defer nc.Close()

	sc, err := stan.Connect(clusterID, clientID, stan.NatsConn(nc)) //stan.NatsURL(URL))
	if err != nil {
		log.Fatalf("Can't connect: %v.\nMake sure a NATS Streaming Server is running at: %s", err, URL)
	}
	defer sc.Close()

	msg := []byte(message)
	// subj := args[0]

	ch := make(chan bool)
	var glock sync.Mutex
	var guid string
	acb := func(lguid string, err error) {
		glock.Lock()
		log.Printf("Received ACK for guid %s\n", lguid)
		defer glock.Unlock()
		if err != nil {
			log.Fatalf("Error in server ack for guid %s: %v\n", lguid, err)
		}
		if lguid != guid {
			log.Fatalf("Expected a matching guid in ack callback, got %s vs %s\n", lguid, guid)
		}
		ch <- true
	}

	if !async {
		err = sc.Publish(subj, msg)
		if err != nil {
			log.Fatalf("Error during publish: %v\n", err)
		}
		log.Printf("Published [%s] : '%s'\n", subj, msg)
	} else {
		glock.Lock()
		guid, err = sc.PublishAsync(subj, msg, acb)
		if err != nil {
			log.Fatalf("Error during async publish: %v\n", err)
		}
		glock.Unlock()
		if guid == "" {
			log.Fatal("Expected non-empty guid to be returned.")
		}
		log.Printf("Published [%s] : '%s' [guid: %s]\n", subj, msg, guid)

		select {
		case <-ch:
			break
		case <-time.After(5 * time.Second):
			log.Fatal("timeout")
		}

	}
}
