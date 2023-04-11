package storage

import (
	"avitoTest/internal/entities"
	"github.com/jmoiron/sqlx"
)

func (s *Storage) CheckOrder(order entities.Order) (bool, error) {
	var result bool
	query := "SELECT exists(SELECT order_id FROM orders WHERE id=$1)"
	row := s.db.QueryRow(query, order.OrderID)
	if err := row.Scan(&result); err != nil {
		return false, err
	}
	return result, nil
}
func (s *Storage) GetOrder(order entities.Order) (entities.Order, error) {
	var orderResult entities.Order
	query := "SELECT * FROM orders WHERE order_id=$1"
	rows, err := s.db.Queryx(query, order.OrderID)
	if err != nil {
		return entities.Order{}, err
	}
	if err := rows.StructScan(&orderResult); err != nil {
		return entities.Order{}, err
	}
	return orderResult, nil
}
func (s *Storage) MakeOrder(tx *sqlx.Tx, order entities.Order) error {
	query := "INSERT INTO orders (user_id, product_id, order_id, price) values ($1, $2, $3, $4)"
	_, err := tx.Exec(query, order.UserID, order.ProductID, order.OrderID, order.Price)
	return err
}
func (s *Storage) RemoveOrder(tx *sqlx.Tx, order entities.Order) error {
	query := "delete from orders where order_id=$1"
	_, err := tx.Exec(query, order.OrderID)
	return err
}
