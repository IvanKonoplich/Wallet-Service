package controllers

import "avitoTest/internal/entities"

type balanceOperations interface {
	BalanceIncrease(user entities.User) error
	TransferOfFunds(transfer entities.Transfer) error
	GetBalance(user entities.User) (float64, error)
}

type reserveOperations interface {
	ReserveFunds(order entities.Order) error
	RevenueApproval(order entities.Order) error
	RevenueDeny(order entities.Order) error
}

type reportOperation interface {
	GetMonthReport(month string) (string, error)
}

type operationsJournal interface {
	GetOperationsListByAmount(user entities.User) (string, error)
	GetOperationsListByDate(user entities.User) (string, error)
}

type useCase interface {
	balanceOperations
	reserveOperations
	reportOperation
	operationsJournal
}

type Controller struct {
	uc useCase
}

func New(uc useCase) *Controller {
	return &Controller{
		uc: uc,
	}
}
