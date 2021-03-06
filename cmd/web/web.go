package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-franky/plain_http/web"
)

func main() {
	port := flag.Int("port", 8080, "port to run web server on")
	flag.Parse()

	srv, err := web.New(
		web.WithLogger(web.BaseLogger),
	)

	if err != nil {
		log.Fatal(err)
	}

	s := &http.Server{
		Handler: srv.Handler,
		Addr:    fmt.Sprintf(":%d", *port),
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	log.Printf("Starting server on port %s\n", s.Addr)
	go func() {
		if err := s.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	graceful(s, web.BaseLogger, 5*time.Second)
}

func graceful(s *http.Server, l web.Logger, timeout time.Duration) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	l.Infof("Shutting down with timeout: %s\n", timeout)
	if err := s.Shutdown(ctx); err != nil {
		l.Infof("Error shutting down: %v\n", err)
	} else {
		l.Infof("Server stopped")
	}
}
