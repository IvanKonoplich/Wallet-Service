package entities

type Order struct {
	OrderID   int     `json:"order_id" binding:"required" db:"order_id"`
	UserID    int     `json:"user_id" binding:"required" db:"user_id"`
	ProductID int     `json:"product_id" binding:"required" db:"product_id"`
	Price     float64 `json:"price" binding:"required" db:"price"`
}
