package main

import (
	"fmt"
	"log"

	"github.com/akrylysov/algnhsa"
	"github.com/go-franky/plain_http/web"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	srv, err := web.New(
		web.WithLogger(web.NoopLogger),
	)
	if err != nil {
		return fmt.Errorf("could not create the server: %w", err)
	}

	algnhsa.ListenAndServe(srv.Handler, nil)
	return nil
}
