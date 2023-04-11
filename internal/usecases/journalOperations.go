package usecases

import (
	"avitoTest/internal/entities"
	"fmt"
)

func (uc *UseCase) GetOperationsListByAmount(user entities.User) ([]entities.OperationsJournalRow, error) {
	result, err := uc.trans.GetOperationsListByAmount(user)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no operations were performed to user:%d. probably user was not created", user.ID)
	}
	return result, nil
}
func (uc *UseCase) GetOperationsListByDate(user entities.User) ([]entities.OperationsJournalRow, error) {
	result, err := uc.trans.GetOperationsListByDate(user)
	if err != nil {
		return nil, err
	}
	if len(result) == 0 {
		return nil, fmt.Errorf("no operations were performed to user:%d. probably user was not created", user.ID)
	}
	return result, nil
}
