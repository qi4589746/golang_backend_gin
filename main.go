package main

import (
	_ "golang_backend_gin/agents"
	"log"
)

func main() {
	router := initRouter()
	err := router.Run(":8000")
	if err != nil {
		log.Fatal("Some error happen while run the Gin server...", err.Error())
	}
}
