package main

import (
	"log"
	"os"

	"go1fl-sprint6-final/internal/server"
)


func main() {
	logger := log.New(os.Stdout, "Server: ", log.LstdFlags)

	myServer := server.CreateHttpRouter(logger)

	logger.Printf("Start server port %s", myServer.Server.Addr)

	err := myServer.Server.ListenAndServe() 
    if err != nil {
       logger.Fatalf("Error server %v", err)
    }
}