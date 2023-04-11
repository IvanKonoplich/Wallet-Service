package usecases

import (
	"avitoTest/internal/entities"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (uc *UseCase) BalanceIncrease(user entities.User) error {
	if user.Balance < 0 {
		return errors.New("incorrect balance request: balance is negative number")
	}
	ok, err := uc.trans.CheckUser(user.ID)
	if err != nil {
		logrus.Errorf("cant check user: %s", err.Error())
		return err
	}
	if ok {
		if err := uc.trans.BalanceIncrease(user); err != nil {
			logrus.Errorf("cant increase balance: %s", err.Error())
			return err
		}
		relevantBalance, err := uc.GetBalance(user)
		if err != nil {
			logrus.Errorf("cant check balance: %s", err.Error())
			return err
		}
		logrus.Infof("Users %d balance has been increased on %f. Now it is: %f", user.ID, user.Balance, relevantBalance)
	} else {
		if err := uc.trans.CreateUser(user); err != nil {
			logrus.Errorf("cant create user: %s", err.Error())
			return err
		}
		relevantBalance, err := uc.GetBalance(user)
		if err != nil {
			logrus.Errorf("cant get balance: %s", err.Error())
			return err
		}
		logrus.Infof("User %d was created. Now his balance is: %f", user.ID, relevantBalance)
	}
	return nil
}

func (uc *UseCase) TransferOfFunds(transfer entities.Transfer) error {
	if transfer.Amount < 0 {
		return errors.New("incorrect transfer amount request: transfer amount is negative number")
	}
	sender, err := uc.trans.CheckUser(transfer.SenderID)
	if err != nil {
		logrus.Errorf("cant check user: %s", err.Error())
		return err
	}
	if sender == false {
		logrus.Infof("cant make transfer of funds becouse sender with id:%d was not created", transfer.SenderID)
		return fmt.Errorf("sender with id:%d was not created", transfer.SenderID)
	}
	recipient, err := uc.trans.CheckUser(transfer.RecipientID)
	if err != nil {
		logrus.Errorf("cant check user: %s", err.Error())
		return err
	}
	if recipient == false {
		logrus.Infof("cant make transfer of funds becouse recipient with id:%d was not created", transfer.RecipientID)
		return fmt.Errorf("sender with id:%d was not created", transfer.RecipientID)
	}
	if err := uc.trans.TransferOfFunds(transfer, false); err != nil {
		logrus.Errorf("cant make transfer: %s", err.Error())
		return err
	}
	logrus.Infof("transfer of funds from user:%d to user:%d was made sucsessfully", transfer.SenderID, transfer.RecipientID)
	return nil
}

func (uc *UseCase) GetBalance(user entities.User) (float64, error) {
	ok, err := uc.trans.CheckUser(user.ID)
	if err != nil {
		logrus.Errorf("cant check user: %s", err.Error())
		return 0, err
	}
	if !ok {
		logrus.Infof("cant get balance becouse user with id:%d was not created", user.ID)
		return 0, fmt.Errorf("user with id:%d was not created", user.ID)
	} else {
		balance, err := uc.trans.GetBalance(user.ID)
		if err != nil {
			logrus.Errorf("cant check balance: %s", err.Error())
			return 0, err
		}
		return balance, nil
	}
}
