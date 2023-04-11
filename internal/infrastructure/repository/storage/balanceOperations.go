package storage

import (
	"avitoTest/internal/entities"
	"github.com/jmoiron/sqlx"
)

func (s *Storage) CheckUser(userID int) (bool, error) {
	var result bool
	query := "SELECT exists(SELECT id FROM users WHERE id=$1)"
	row := s.db.QueryRow(query, userID)
	if err := row.Scan(&result); err != nil {
		return false, err
	}
	return result, nil
}
func (s *Storage) CreateUser(tx *sqlx.Tx, user entities.User) error {
	query := "INSERT INTO users (id, balance) values ($1, $2)"
	_, err := tx.Exec(query, user.ID, user.Balance)
	return err

}
func (s *Storage) BalanceIncrease(tx *sqlx.Tx, userID int, amount float64) error {
	query := "UPDATE users SET balance=balance+$1 WHERE id=$2"
	_, err := tx.Exec(query, amount, userID)
	return err
}
func (s *Storage) BalanceDecrease(tx *sqlx.Tx, userID int, amount float64) error {
	query := "UPDATE users SET balance=balance-$1 WHERE id=$2"
	_, err := tx.Exec(query, amount, userID)
	return err
}
func (s *Storage) GetBalance(userid int) (float64, error) {
	var balance float64
	query := "SELECT balance FROM users WHERE id=$1"
	row := s.db.QueryRow(query, userid)
	if err := row.Scan(&balance); err != nil {
		return 0, err
	}
	return balance, nil
}
func (s *Storage) CheckBalance(tx *sqlx.Tx, userID int) (float64, error) {
	var balance float64
	query := "SELECT balance FROM users WHERE id=$1 FOR UPDATE "
	row := tx.QueryRow(query, userID)
	if err := row.Scan(&balance); err != nil {
		return 0, err
	}
	return balance, nil
}
