package transactionsMaker

import "avitoTest/internal/entities"

func (tm *TransactionMaker) GetOperationsListByAmount(user entities.User) ([]entities.OperationsJournalRow, error) {
	return tm.repos.GetOperationsListByAmount(user)
}
func (tm *TransactionMaker) GetOperationsListByDate(user entities.User) ([]entities.OperationsJournalRow, error) {
	return tm.repos.GetOperationsListByDate(user)
}
