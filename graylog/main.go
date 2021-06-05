package main

import (
	"time"
)

func main() {
	// Initialize a new graylog client with TLS
	g, err := NewGraylog(Endpoint{
		Transport: "tcp",
		Address:   "127.0.0.1",
		Port:      12201,
	})
	if err != nil {
		panic(err)
	}

	// Send a message
	err = g.Send(Message{
		Version:      "1.1",
		Host:         "127.0.0.1",
		ShortMessage: "Sample test",
		FullMessage:  "Stacktrace",
		Timestamp:    time.Now().Unix(),
		Level:        1,
		Extra: map[string]string{
			"MY-EXTRA-FIELD": "extra_value",
		},
	})
	if err != nil {
		panic(err)
	}

	// Close the graylog connection
	if err := g.Close(); err != nil {
		panic(err)
	}
}
