package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)


func main() {
	logger := log.New(os.Stdout, "Server: ", log.LstdFlags)

	myServer := server.CreateHttpRouter(logger)
	err := myServer.Server.ListenAndServe() 
    if err != nil {
       logger.Fatalf("Error server %v", err)
    }
}