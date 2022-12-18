package main

import (
	"context"
	"flag"
	"graph/internal/graph"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

var _port = flag.Int("p", 8080, "server port")

func main() {
	r := mux.NewRouter()
	handler := &graph.Handler{}

	r.Handle("/convertMatrix", handler).Methods("POST")
	s := http.Server{
		Addr:    net.JoinHostPort("", strconv.Itoa(*_port)),
		Handler: r,
	}
	go func() {
		log.Printf("server started on port: %d", *_port)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
		log.Printf("server was shuted down")
	}()

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	<-exit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
