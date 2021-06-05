package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	defaulthttp "net/http"
	"time"

	"github.com/arangodb/go-driver"
	"github.com/arangodb/go-driver/http"
)

var (
	fUsr  = flag.String("name", "root", "access user")
	fPwd  = flag.String("password", "somepassword", "access password")
	fAddr = flag.String("addr", "http://127.0.0.1:8529", "access addrss")
)

const (
	MaxIdleConns        int = 60000
	MaxIdleConnsPerHost int = 10000
	IdleConnTimeout     int = 150
	RequestTimeout      int = 30
)

func main() {
	flag.Parse()

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
		Endpoints: []string{*fAddr},
		Transport: cTransport,
	})
	if err != nil {
		panic(err)
	}

	c, err := driver.NewClient(driver.ClientConfig{
		Connection:     conn,
		Authentication: driver.BasicAuthentication(*fUsr, *fPwd),
	})

	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	db, err := c.Database(ctx, "_system")
	if err != nil {
		panic(err)
	}

	found, err := db.CollectionExists(ctx, "AdminAccounts")
	if err != nil {
		panic(err)
	}
	fmt.Println(found)
}
