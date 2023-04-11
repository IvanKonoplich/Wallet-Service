package storage

import "github.com/jmoiron/sqlx"

func (s *Storage) CreateTx() (tx *sqlx.Tx, err error) {
	return s.db.Beginx()
}
