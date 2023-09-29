package main

import (
	"log"
	"net/http"

	"github.com/sliseev/websocket_service/back/pkg/ws"
)

func serveWs(controller *ws.Controller, w http.ResponseWriter, r *http.Request) {
	conn, err := ws.Upgrade(w, r)
	if err != nil {
		log.Printf("serveWs: %v", err)
		return
	}

	client := &ws.Client{
		Conn:       conn,
		Controller: controller,
	}

	if err := client.ReadName(); err != nil {
		log.Printf("serveWs: read name failed: %v", err)
		return
	}

	controller.Register <- client
	client.ProcessMessages()
}

func setupRoutes() {
	controller := ws.NewController()
	go controller.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(controller, w, r)
	})
}

func main() {
	log.Println("Chat App v0.01")
	setupRoutes()
	http.ListenAndServe(":8080", nil)
}
