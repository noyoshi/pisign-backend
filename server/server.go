// Package server much of the code comes from https://tutorialedge.net/projects/chat-system-in-go-and-react/part-4-handling-multiple-clients/
package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/pisign/pisign-backend/widget"
)

func serveWs(pool *widget.Pool, w http.ResponseWriter, r *http.Request) {
	log.Println("Websocket endpoing hit!")
	conn, err := upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	apiName := r.FormValue("api")

	widget.Create(apiName, conn, pool)
}

func setupRoutes() {
	pool := widget.NewPool()
	go pool.Start()
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})
}

// StartLocalServer creates a new server on localhost
func StartLocalServer(port int) {
	addr := fmt.Sprintf("localhost:%v", port)
	log.Printf("Running server at %v\n", addr)
	setupRoutes()
	http.ListenAndServe(addr, nil)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade->upgrader:", err)
		return ws, err
	}
	return ws, nil
}
