package main

import (
	"fmt"
	"github.com/destinerikanb/go-react-chat-app/pkg/handlers"
	"net/http"
)

func main() {
	fmt.Println("Chat App v1.0.0")
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handlers.SetupRoutes(),
	}

	server.ListenAndServe()

}
