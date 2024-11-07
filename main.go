package main

import (
	"device-service/router"
	"fmt"
	"log"
	"net/http"

	"device-service/configs"
	"device-service/db"

	"github.com/gorilla/mux"
)

func main() {
	config, err := configs.LoadConfig("configs/config.json")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	db, err := db.NewDB(*config)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer db.Close()

	r := mux.NewRouter()
	router.InitializeRoutes(r, &db)
	addressPort := config.Port

	log.Println("Starting server on :", addressPort)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", addressPort), r); err != nil {
		log.Fatal(err)
	}
}
