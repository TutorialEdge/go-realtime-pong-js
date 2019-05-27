package router

import (
	"fmt"
	"net/http"
)

type Router interface {
	SetupRoutes()
}

type WSRouter struct{}

func (r *WSRouter) SetupRoutes() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home Page!")
	})

	http.HandleFunc("/connect", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Connecting...")
	})
}
