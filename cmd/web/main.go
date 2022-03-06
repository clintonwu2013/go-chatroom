package main

import (
	"go-websocket/internal/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	go handlers.ListenToWsChannel()

	log.Println("Starting go websocket server on port 8080 !!!")
	_ = http.ListenAndServe(":8080", mux)
}
