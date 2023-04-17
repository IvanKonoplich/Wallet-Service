package storage

import (
	"avitoTest/internal/entities"
	"encoding/csv"
	"fmt"
	"github.com/jmoiron/sqlx"
	"os"
	"time"
)

func (s *Storage) GetMonthReport(month string) (string, error) {

	query := "SELECT sum(balance) FROM report WHERE month=$1 GROUP BY product_id"
	rows, err := s.db.Queryx(query, month)

	file, err := os.Create("report:" + time.Now().String() + fmt.Sprintf("-month:%s", month) + ".csv")
	if err != nil {
		return "", err
	}
	writer := csv.NewWriter(file)
	for rows.Next() {
		var reportRow entities.MonthReportRow
		err = rows.StructScan(&reportRow)
		if err != nil {
			return "", err
		}
		writer.Write([]string{fmt.Sprintf("%d:%f", reportRow.ProductID, reportRow.Balance)})
	}
	writer.Flush()
	return file.Name(), err
}
func (s *Storage) UpdateMonthReport(tx *sqlx.Tx, order entities.Order, month string) error {
	query := "INSERT INTO report (product_id, balance, month) values ($1, $2, $3)"
	_, err := tx.Exec(query, order.ProductID, order.Price, month)
	return err
}
