package main

import (
	"fmt"

	"github.com/TutorialEdge/go-realtime-pong-js/router"
)


func main() {
	fmt.Println("Network Pong Backend!")
	
	var router router.WSRouter
	router.SetupRoutes()
	router.Start()
}