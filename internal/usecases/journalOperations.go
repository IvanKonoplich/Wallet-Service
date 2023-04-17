package usecases

import (
	"avitoTest/internal/entities"
	"fmt"
)

func (uc *UseCase) GetOperationsListByAmount(user entities.User) (string, error) {
	result, err := uc.trans.GetOperationsListByAmount(user)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", fmt.Errorf("no operations were performed to user:%d. probably user was not created", user.ID)
	}
	resultString := ""
	for _, i := range result {
		resultString += i.Message
		resultString += "\n"
	}
	return resultString, nil
}
func (uc *UseCase) GetOperationsListByDate(user entities.User) (string, error) {
	result, err := uc.trans.GetOperationsListByDate(user)
	if err != nil {
		return "", err
	}
	if len(result) == 0 {
		return "", fmt.Errorf("no operations were performed to user:%d. probably user was not created", user.ID)
	}
	resultString := ""
	for _, i := range result {
		resultString += i.Message
		resultString += "\n"
	}
	return resultString, nil
}
