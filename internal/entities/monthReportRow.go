package entities

type MonthReportRow struct {
	ProductID int     `json:"product_id" db:"product_id"`
	Balance   float64 `json:"balance" db:"sum"`
}
