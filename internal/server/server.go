package server

import (
	"log"
	"net/http"
	"time"

	"go1fl-sprint6-final/internal/handlers"
)

type Server struct{
	Logger *log.Logger
	Server *http.Server
}

func CreateHttpRouter(log *log.Logger) *Server{
mux := http.NewServeMux()


	mux.HandleFunc("/", handlers.GetHtmlFormat)
	mux.HandleFunc("/upload", handlers.UploadFile)
	
	httpServer := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ErrorLog: log,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	return &Server{Logger: log, Server: httpServer}
}

