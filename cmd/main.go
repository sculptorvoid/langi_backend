package main

import (
	langi "github.com/sculptorvoid/langi_backend"
	"log"
)

func main() {
	server := new(langi.Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("Error starting http server: %s", err.Error())
	}
}
