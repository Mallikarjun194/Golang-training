package main

import (
	"blog-project/config"
	"blog-project/database"
	"blog-project/routes"
	"fmt"
	"net/http"
	"os"
)

func main() {

	db := database.OpenDBConn()
	routerObj := routes.Router{DatabaseObj: db}

	router := routerObj.SetupRouter()
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
