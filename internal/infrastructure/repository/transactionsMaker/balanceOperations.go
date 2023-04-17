package transactionsMaker

import (
	"avitoTest/internal/entities"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func (tm *TransactionMaker) CheckUser(userID int) (bool, error) {
	return tm.repos.CheckUser(userID)
}
func (tm *TransactionMaker) CreateUser(user entities.User) error {
	tx, err := tm.repos.CreateTx()
	if err != nil {
		return err
	}
	err = tm.repos.CreateUser(tx, user)
	if err != nil {
		tx.Rollback()
		return err
	}
	journalUpdate := entities.OperationsJournalRow{
		UserID:  user.ID,
		Amount:  user.Balance,
		Date:    time.Now(),
		Message: fmt.Sprintf("Account with ID:%d was created, balance:%f. Date:%s", user.ID, user.Balance, time.Now().String()),
	}
	err = tm.repos.UpdateOperationsJournal(tx, journalUpdate)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (tm *TransactionMaker) BalanceIncrease(user entities.User) error {
	tx, err := tm.repos.CreateTx()
	if err != nil {
		return err
	}
	err = tm.repos.BalanceIncrease(tx, user.ID, user.Balance)
	if err != nil {
		tx.Rollback()
		return err
	}
	relevantBalance, err := tm.GetBalance(user.ID)
	if err != nil {
		tx.Rollback()
		return err
	}
	journalUpdate := entities.OperationsJournalRow{
		UserID:  user.ID,
		Amount:  user.Balance,
		Date:    time.Now(),
		Message: fmt.Sprintf("Account:%d balance was increased on:%f. Now it is:%f. Date:%s", user.ID, user.Balance, relevantBalance, time.Now().String()),
	}
	err = tm.repos.UpdateOperationsJournal(tx, journalUpdate)
	if err != nil {
		tx.Rollback()
		return err
	}
	logrus.Infof("operations journal sucsessfully updated:%s", journalUpdate.Message)
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (tm *TransactionMaker) TransferOfFunds(transfer entities.Transfer, mustCheckIfEnough bool) error {
	tx, err := tm.repos.CreateTx()
	if err != nil {
		return err
	}
	senderBalance, err := tm.repos.CheckBalance(tx, transfer.SenderID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if senderBalance < transfer.Amount {
		return fmt.Errorf("cant make transfer of funds because sender have not enough funds. balance:%f transfer amount:%f", senderBalance, transfer.Amount)
	}
	err = tm.repos.BalanceDecrease(tx, transfer.SenderID, transfer.Amount)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tm.repos.BalanceIncrease(tx, transfer.RecipientID, transfer.Amount)
	if err != nil {
		tx.Rollback()
		return err
	}
	relevantSenderBalance, err := tm.GetBalance(transfer.SenderID)
	if err != nil {
		tx.Rollback()
		return err
	}
	relevantSenderBalance -= transfer.Amount
	relevantRecipientBalance, err := tm.GetBalance(transfer.RecipientID)
	if err != nil {
		tx.Rollback()
		return err
	}
	relevantRecipientBalance += transfer.Amount
	journalUpdateSender := entities.OperationsJournalRow{
		UserID:  transfer.SenderID,
		Amount:  transfer.Amount,
		Date:    time.Now(),
		Message: fmt.Sprintf("Account:%d balance was decreased on:%f, becouse of transfer to user:%d .Now it is:%f. Date:%s", transfer.SenderID, transfer.Amount, transfer.RecipientID, relevantSenderBalance, time.Now().String()),
	}
	journalUpdateRecipient := entities.OperationsJournalRow{
		UserID:  transfer.RecipientID,
		Amount:  transfer.Amount,
		Date:    time.Now(),
		Message: fmt.Sprintf("Account:%d balance was increased on:%f, becouse of transfer from user:%d .Now it is:%f. Date:%s", transfer.RecipientID, transfer.Amount, transfer.SenderID, relevantRecipientBalance, time.Now().String()),
	}
	err = tm.repos.UpdateOperationsJournal(tx, journalUpdateSender)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tm.repos.UpdateOperationsJournal(tx, journalUpdateRecipient)
	if err != nil {
		tx.Rollback()
		return err
	}
	logrus.Infof("operations journal sucsessfully updated:%s, %s", journalUpdateSender.Message, journalUpdateRecipient.Message)
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (tm *TransactionMaker) GetBalance(userid int) (float64, error) {
	return tm.repos.GetBalance(userid)
}
