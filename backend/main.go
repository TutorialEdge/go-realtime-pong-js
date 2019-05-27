package main

import (
	"fmt"
	"github.com/TutorialEdge/go-realtime-pong-js/backend/router"
)


func main() {
	fmt.Println("Network Pong Backend!")
	router.SetupRoutes()
	
}