package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ganinugroho/belajar/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api ", log.LstdFlags)
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodBye(l)
	ph := handlers.NewProducts(l)
	sm := http.NewServeMux()
	sm.Handle("/hello", hh)
	sm.Handle("/bye", gh)
	sm.Handle("/product/", ph)

	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received termintal, graceful shutdown", sig)

	tc, err := context.WithTimeout(context.Background(), 30*time.Second)

	if err != nil {
		s.Shutdown(tc)
	} else {
		l.Fatal(err)
	}
}
