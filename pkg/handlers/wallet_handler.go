package handlers

import (
	"encoding/json"
	"github.com/google/uuid"
	"log"
	"net/http"
	"wallet-service/pkg/config"
	"wallet-service/pkg/db"
	"wallet-service/pkg/models"
)

type WalletRequest struct {
	WalletID  uuid.UUID `json:"walletId"`
	Operation string    `json:"operationType"`
	Amount    float64   `json:"amount"`
}

// WalletHandler - обработчик запроса для обновления баланса
func WalletHandler(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Failed to load config:", err)
		http.Error(w, "Failed to load config", http.StatusInternalServerError)
		return
	}

	dbConn, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	var req WalletRequest
	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		log.Println("Invalid request body:", err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if req.Operation == "DEPOSIT" {
		err = db.UpdateBalance(dbConn, req.WalletID.String(), req.Amount)
	} else if req.Operation == "WITHDRAW" {
		err = db.UpdateBalance(dbConn, req.WalletID.String(), -req.Amount)
	} else {
		http.Error(w, "Invalid operation type", http.StatusBadRequest)
		return
	}

	if err != nil {
		log.Println("Failed to update balance:", err)
		http.Error(w, "Failed to update balance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Balance updated"))
}

// GetWalletBalance - обработчик запроса для получения баланса
func GetWalletBalance(w http.ResponseWriter, r *http.Request) {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println("Failed to load config:", err)
		http.Error(w, "Failed to load config", http.StatusInternalServerError)
		return
	}

	dbConn, err := db.ConnectToDB(cfg)
	if err != nil {
		log.Println("Failed to connect to database:", err)
		http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
		return
	}
	defer dbConn.Close()

	walletID := r.URL.Path[len("/api/v1/wallets/"):]

	balance, err := db.GetBalance(dbConn, walletID)
	if err != nil {
		log.Println("Failed to get balance:", err)
		http.Error(w, "Failed to get balance", http.StatusInternalServerError)
		return
	}

	res := models.Wallet{
		WalletID: uuid.MustParse(walletID),
		Balance:  balance,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}
