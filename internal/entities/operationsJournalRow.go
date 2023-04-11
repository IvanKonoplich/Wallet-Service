package entities

import "time"

type OperationsJournalRow struct {
	UserID  int       `json:"user_id" binding:"required" db:"user_id"`
	Amount  float64   `json:"amount" binding:"required" db:"amount"`
	Date    time.Time `json:"date" binding:"required" db:"date"`
	Message string    `json:"message" binding:"required" db:"message" `
}
