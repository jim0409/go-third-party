package main

import (
	"context"
	"net"
	defaulthttp "net/http"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

const (
	MaxIdleConns        int    = 60000
	MaxIdleConnsPerHost int    = 10000
	IdleConnTimeout     int    = 150
	RequestTimeout      int    = 30
	OperatingDatabase   string = "_system"
	// OperatingDatabase   string = "Database"
)

func NewArangoDB(addr string) (driver.Database, error) {
	cTransport := &defaulthttp.Transport{
		Proxy: defaulthttp.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout:   time.Duration(RequestTimeout) * time.Second,
			KeepAlive: time.Duration(RequestTimeout) * time.Second,
			DualStack: true,
		}).DialContext,
		MaxIdleConns:          MaxIdleConns,
		MaxIdleConnsPerHost:   MaxIdleConnsPerHost,
		IdleConnTimeout:       time.Duration(IdleConnTimeout) * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	conn, err := http.NewConnection(http.ConnectionConfig{
		Endpoints: []string{addr},
		Transport: cTransport,
	})
	if err != nil {
		return nil, err
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication("root", "yourpassword"),
	})
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	db, err := c.Database(ctx, OperatingDatabase)

	return db, err
}
