package main

import (
	"log"
	"net/http"
	"websocket-chat-app/internal/server"
)

func main() {
	server.SetupRoutes()

	log.Println("The server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
