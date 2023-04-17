package transactionsMaker

import (
	"avitoTest/internal/entities"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"
)

func (tm *TransactionMaker) CheckOrder(order entities.Order) (bool, error) {
	return tm.repos.CheckOrder(order)
}
func (tm *TransactionMaker) GetOrder(order entities.Order) (entities.Order, error) {
	return tm.repos.GetOrder(order)
}
func (tm *TransactionMaker) ReserveFunds(order entities.Order) error {
	tx, err := tm.repos.CreateTx()
	if err != nil {
		return err
	}
	relevantBalance, err := tm.repos.CheckBalance(tx, order.UserID)
	if err != nil {
		tx.Rollback()
		return err
	}
	if relevantBalance < order.Price {
		tx.Rollback()
		return fmt.Errorf("cant reserve funds because user have not enough money. balance:%f, price:%f", relevantBalance, order.Price)
	}
	err = tm.repos.BalanceDecrease(tx, order.UserID, order.Price)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tm.repos.MakeOrder(tx, order)
	if err != nil {
		tx.Rollback()
		return err
	}
	journalUpdate := entities.OperationsJournalRow{
		UserID:  order.UserID,
		Amount:  order.Price,
		Date:    time.Now(),
		Message: fmt.Sprintf("account:%d funds were reserved for order:%d. Amount:%f. Product:%d. Date:%s", order.UserID, order.OrderID, order.Price, order.ProductID, fmt.Sprint(time.Now().Date())+""+fmt.Sprint(time.Now().Clock())),
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
func (tm *TransactionMaker) RevenueApproval(order entities.Order) error {
	tx, err := tm.repos.CreateTx()
	if err != nil {
		return err
	}
	err = tm.repos.RemoveOrder(tx, order)
	if err != nil {
		tx.Rollback()
		return err
	}
	logrus.Infof("company balance increased")
	err = tm.repos.UpdateMonthReport(tx, order, time.Now().Month().String())
	if err != nil {
		tx.Rollback()
		return err
	}
	logrus.Infof("mounth report sucsessfully updated")
	journalUpdate := entities.OperationsJournalRow{
		UserID:  order.UserID,
		Amount:  order.Price,
		Date:    time.Now(),
		Message: fmt.Sprintf("account:%d order:%d funds approved. Amount:%f. Product:%d. Date:%s", order.UserID, order.OrderID, order.Price, order.ProductID, fmt.Sprint(time.Now().Date())+""+fmt.Sprint(time.Now().Clock())),
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
func (tm *TransactionMaker) RevenueDeny(order entities.Order) error {
	tx, err := tm.repos.CreateTx()
	if err != nil {
		return err
	}
	err = tm.repos.RemoveOrder(tx, order)
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tm.repos.BalanceIncrease(tx, order.UserID, order.Price)
	if err != nil {
		tx.Rollback()
		return err
	}

	journalUpdate := entities.OperationsJournalRow{
		UserID:  order.UserID,
		Amount:  order.Price,
		Date:    time.Now(),
		Message: fmt.Sprintf("account:%d order:%d funds denyed. Amount:%f. Product:%d. Date:%s", order.UserID, order.OrderID, order.Price, order.ProductID, fmt.Sprint(time.Now().Date())+""+fmt.Sprint(time.Now().Clock())),
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
