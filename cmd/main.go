package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"wallet-service/pkg/config"
	"wallet-service/pkg/handlers"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	r := mux.NewRouter()

	r.HandleFunc("/api/v1/wallet", handlers.WalletHandler).Methods("POST")
	r.HandleFunc("/api/v1/wallets/{walletID}", handlers.GetWalletBalance).Methods("GET")

	log.Printf("Starting server on port %s...\n", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, r); err != nil {
		log.Fatal("Error starting server:", err)
	}
}
