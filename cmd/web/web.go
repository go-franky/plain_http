package main

import (
	"flag"
	"log"

	"github.com/go-franky/plain_http/web"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.New()
	port := flag.Int("port", 8080, "port to run web server on")
	flag.Parse()

	srv, err := web.New(web.WithLogger(logger), web.WithPort(*port))
	if err != nil {
		log.Fatal(err)
	}
	srv.Start()
}
