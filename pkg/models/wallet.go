package models

import "github.com/google/uuid"

type Wallet struct {
	WalletID uuid.UUID `json:"walletId"`
	Balance  float64   `json:"balance"`
}
