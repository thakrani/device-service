package main

import (
	"device-service/router"
	"fmt"
	"log"
	"net/http"

	"device-service/configs"

	"github.com/gorilla/mux"
)

func main() {
	config, err := configs.LoadConfig("configs/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	r := mux.NewRouter()
	router.InitializeRoutes(r)
	addressPort := config.Port

	log.Println("Starting server on :", addressPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", addressPort), r); err != nil {
		log.Fatal(err)
	}
}
