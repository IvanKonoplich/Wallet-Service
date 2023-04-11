package transactionsMaker

import (
	"avitoTest/internal/entities"
	"github.com/jmoiron/sqlx"
)

type transactionsOperations interface {
	CreateTx() (tx *sqlx.Tx, err error)
}
type balanceOperations interface {
	CheckUser(userID int) (bool, error)
	CreateUser(tx *sqlx.Tx, user entities.User) error
	BalanceIncrease(tx *sqlx.Tx, userID int, amount float64) error
	BalanceDecrease(tx *sqlx.Tx, userID int, amount float64) error
	GetBalance(userid int) (float64, error)
	CheckBalance(tx *sqlx.Tx, userID int) (float64, error)
}

type reserveOperations interface {
	CheckOrder(order entities.Order) (bool, error)
	GetOrder(order entities.Order) (entities.Order, error)
	MakeOrder(tx *sqlx.Tx, order entities.Order) error
	RemoveOrder(tx *sqlx.Tx, order entities.Order) error
}

type reportOperations interface {
	GetMonthReport(month string) (string, error)
	UpdateMonthReport(tx *sqlx.Tx, order entities.Order, month string) error
}

type operationsJournal interface {
	UpdateOperationsJournal(tx *sqlx.Tx, row entities.OperationsJournalRow) error
	GetOperationsListByAmount(user entities.User) ([]entities.OperationsJournalRow, error)
	GetOperationsListByDate(user entities.User) ([]entities.OperationsJournalRow, error)
}
type repos interface {
	balanceOperations
	reserveOperations
	reportOperations
	operationsJournal
	transactionsOperations
}

type TransactionMaker struct {
	repos repos
}

func NewTransactionMaker(store repos) *TransactionMaker {
	return &TransactionMaker{repos: store}
}
