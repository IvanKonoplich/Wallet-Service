package storage

import (
	"avitoTest/internal/entities"
	"github.com/jmoiron/sqlx"
)

func (s *Storage) UpdateOperationsJournal(tx *sqlx.Tx, row entities.OperationsJournalRow) error {
	query := "INSERT INTO operations_journal (user_id, amount, date, message) values ($1, $2, $3, $4)"
	_, err := tx.Exec(query, row.UserID, row.Amount, row.Date, row.Message)
	return err
}
func (s *Storage) GetOperationsListByAmount(user entities.User) ([]entities.OperationsJournalRow, error) {
	var result []entities.OperationsJournalRow
	query := "SELECT * FROM operations_journal order by amount"
	rows, err := s.db.Queryx(query)
	for rows.Next() {
		var row entities.OperationsJournalRow
		err = rows.StructScan(&row)
		if err != nil {
			return []entities.OperationsJournalRow{}, err
		}
		result = append(result, row)
	}
	return result, nil
}
func (s *Storage) GetOperationsListByDate(user entities.User) ([]entities.OperationsJournalRow, error) {
	var result []entities.OperationsJournalRow
	query := "SELECT * FROM operations_journal order by date"
	rows, err := s.db.Queryx(query)
	for rows.Next() {
		var row entities.OperationsJournalRow
		err = rows.StructScan(&row)
		if err != nil {
			return nil, err
		}
		result = append(result, row)
	}
	return result, nil
}
