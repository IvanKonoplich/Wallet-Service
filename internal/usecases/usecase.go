package usecases

import (
	"avitoTest/internal/entities"
)

//go:generate mockgen -source=usecase.go -destination=mocks/usecaseMock.go

type balanceOperations interface {
	CheckUser(userID int) (bool, error)
	CreateUser(user entities.User) error
	BalanceIncrease(user entities.User) error
	TransferOfFunds(transfer entities.Transfer, mustCheckIfEnough bool) error
	GetBalance(userid int) (float64, error)
}

type reserveOperations interface {
	CheckOrder(order entities.Order) (bool, error)
	GetOrder(order entities.Order) (entities.Order, error)
	ReserveFunds(order entities.Order) error
	RevenueApproval(order entities.Order) error
	RevenueDeny(order entities.Order) error
}

type reportOperations interface {
	GetMonthReport(month string) (string, error)
}
type operationsJournal interface {
	GetOperationsListByAmount(user entities.User) ([]entities.OperationsJournalRow, error)
	GetOperationsListByDate(user entities.User) ([]entities.OperationsJournalRow, error)
}
type transactionMaker interface {
	balanceOperations
	reserveOperations
	reportOperations
	operationsJournal
}

type UseCase struct {
	trans transactionMaker
}

func New(trans transactionMaker) *UseCase {
	return &UseCase{
		trans: trans,
	}
}
