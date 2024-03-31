package handlers

import (
	web_socket "github.com/destinerikanb/go-react-chat-app/pkg/web-socket"
	"github.com/gorilla/mux"
	"net/http"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	pool := web_socket.NewPool()
	go pool.Start()

	// ws endpoint
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})

	return router
}
