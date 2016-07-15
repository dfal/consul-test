package main

import (
	"fmt"
	"github.com/op/go-logging"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var log = logging.MustGetLogger("service2")

func main() {

	log.Info("Starting Service 2 ...")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		for _ = range c {
			log.Info("Stopping the service...")
			os.Exit(0)
		}
	}()

	go runPolling()
	runCheckHealthEndpoint()
}

func runPolling() {
	for i := 0; ; i++ {
		response, err := http.Get("http://172.20.20.11:9090/Dima")
		if err != nil {
			log.Error(err)
		} else {
			defer response.Body.Close()
			responseBody, err := ioutil.ReadAll(response.Body)

			if err != nil {
				log.Error(err)
			} else {
				log.Notice(string(responseBody))
			}
		}

		log.Info("Sleeping...")

		time.Sleep(5 * time.Second)
	}
}

func runCheckHealthEndpoint() {
	http.HandleFunc("/check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Service2 - OK")
	})

	http.ListenAndServe(":9091", nil)
}
