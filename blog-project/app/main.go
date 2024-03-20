package main

import (
	"blog-project/config"
	"blog-project/routes"
	"fmt"
	"net/http"
	"os"
)

func main() {
	router := routes.SetupRouter()
	port := os.Getenv("PORT")
	if port == "" {
		port = config.DEFAULT_PORT
	}
	fmt.Println("Starting a service in port!!!", port)
	server := &http.Server{
		Addr:    "localhost:" + port,
		Handler: router,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic("Failed to start a server!!!!!")
	}

}
