package main

import (
	"fmt"
	"github.com/stretchr/goweb"
	"github.com/stretchr/goweb/context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

const (
	Address string = ":9090"
)

func main() {

	mapRoutes()

	log.Print("Starting Service 1 ...")

	s := &http.Server{
		Addr:    Address,
		Handler: goweb.DefaultHttpHandler(),
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	listener, listenErr := net.Listen("tcp", Address)

	log.Printf("  visit: %s", Address)

	if listenErr != nil {
		log.Fatalf("Could not listen: %s", listenErr)
	}

	go func() {
		for _ = range c {
			log.Print("Stopping the server...")
			listener.Close()
			os.Exit(0)
		}
	}()

	log.Fatalf("Error in Serve: %s", s.Serve(listener))
}

func mapRoutes() {
	goweb.Map("/check", func(c context.Context) error {
		return goweb.API.Respond(c, 200, "Service1 - OK", nil)
	})

	goweb.Map("/[name]", func(c context.Context) error {
		if c.PathParams().Has("name") {
			return goweb.API.Respond(c, 200, fmt.Sprintf("Hello, %s", c.PathParams().Get("name")), nil)
		} else {
			return goweb.API.Respond(c, 200, "Hello!", nil)
		}
	})
}
