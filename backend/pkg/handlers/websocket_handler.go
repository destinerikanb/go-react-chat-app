package handlers

import (
	"fmt"
	"github.com/destinerikanb/go-react-chat-app/pkg/web-socket"
	"net/http"
	"strconv"
)

// define ws handler
func serveWS(pool *web_socket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("Websockect endpoint hit")

	ws, err := web_socket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	counter := 1

	client := &web_socket.Client{
		Name: "user " + strconv.Itoa(counter),
		Conn: ws,
		Pool: pool,
	}

	pool.Register <- client
	counter++
	client.Read()
}
