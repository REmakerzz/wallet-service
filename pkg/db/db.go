package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"wallet-service/pkg/config"
)

// ConnectToDB - подключение к базе данных
func ConnectToDB(cfg *config.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

// UpdateBalance - обновление баланса кошелька
func UpdateBalance(db *sql.DB, walletID string, amount float64) error {
	query := `
		UPDATE wallets
		SET balance = balance + $1
		WHERE wallet_id = $2
	`
	_, err := db.Exec(query, amount, walletID)
	return err
}

// GetBalance - получение баланса кошелька
func GetBalance(db *sql.DB, walletID string) (float64, error) {
	var balance float64
	query := `SELECT balance FROM wallets WHERE wallet_id = $1`
	err := db.QueryRow(query, walletID).Scan(&balance)
	if err != nil {
		return 0, err
	}
	return balance, nil
}
