package main

import (
	"cryptoapi/pkg/httpserver"
	"log"
)

func main() {
	srv := httpserver.NewServer()
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
