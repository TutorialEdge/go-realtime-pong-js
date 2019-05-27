package router

import (
	"fmt"
	"net/http"

	"github.com/TutorialEdge/go-realtime-pong-js/websocket"
)

type Router interface {
	SetupRoutes()
}

type WSRouter struct{}

func (r *WSRouter) SetupRoutes() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home Page!")
	})

	http.HandleFunc("/connect", websocket.Connect)
}

func (r *WSRouter) Start() {
	http.ListenAndServe(":10000", nil)
}
