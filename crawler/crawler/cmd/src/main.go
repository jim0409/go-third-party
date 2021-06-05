package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jimweng/crawler/crawler/config"
	"github.com/jimweng/crawler/crawler/utils"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	action = flag.String("act", "", "please check help doc")
	addr   = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")
)

var cfg *config.CrawlerConfig

func main() {
	flag.Parse()

	cfg = config.NewConfig(*action)
	if cfg == nil {
		os.Exit(1)
	}
	var c chan string = make(chan string)

	go matric(c, 5)

	go agent(c, 5)
	// collect()
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Printf("%v\n", sig)

		done <- true
	}()

	<-done
}

type HelloHandler struct{}

func matric(c chan string, t int) {
	http.Handle("/metrics", promhttp.Handler())

	helloHandler := HelloHandler{}
	http.Handle("/_health", helloHandler)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

/*
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
*/
func (h HelloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("{status:'OK'}")
}

func collect() {
	var points *[]*utils.PKGContent

	for _, j := range cfg.InputPlugins {
		if pts, err := j.Gather(); err != nil {
			log.Fatal("%v\n", err)
		} else {
			points = pts.(*[]*utils.PKGContent)
		}
	}

	for _, j := range cfg.OutputPlugins {
		if err := j.Write(points); err != nil {
			log.Fatal("%v\n", err)
		}

	}
}

func agent(c chan string, t int) {
	for {
		collect()
		time.Sleep(time.Second * time.Duration(t))
	}
}

func printc(c chan string, t int) {

}

func init() {
}
