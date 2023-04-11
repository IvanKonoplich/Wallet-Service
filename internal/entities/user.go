package entities

type User struct {
	ID      int     `json:"id" binding:"required"`
	Balance float64 `json:"amount"`
}
