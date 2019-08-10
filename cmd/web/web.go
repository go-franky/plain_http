package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

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
	fmt.Printf("Starting server on port %d\n", *port)
	log.Fatal(http.ListenAndServe(":"+fmt.Sprintf("%d", *port), srv.Handler))
}
