package usecases

import (
	"avitoTest/internal/entities"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

func (uc *UseCase) ReserveFunds(order entities.Order) error {
	if order.Price < 0 {
		return errors.New("incorrect price: price is negative number")
	}
	ok, err := uc.trans.CheckUser(order.UserID)
	if err != nil {
		logrus.Errorf("cant check user: %s", err.Error())
		return err
	}
	if !ok {
		logrus.Infof("cant reserve funds becouse user with id:%d was not created", order.UserID)
		return fmt.Errorf("user with id:%d was not created", order.UserID)
	} else {
		ok, err := uc.trans.CheckOrder(order)
		if ok {
			return fmt.Errorf("order with id:%d already exists", order.OrderID)
		}
		err = uc.trans.ReserveFunds(order)
		if err != nil {
			logrus.Errorf("cant reserve funds: %s", err.Error())
			return err
		}
		logrus.Infof("funds from user:%d were sucsessfuly reserved on order:%d", order.UserID, order.OrderID)
		return nil
	}
}

func (uc *UseCase) RevenueApproval(order entities.Order) error {
	ok, err := uc.trans.CheckOrder(order)
	if err != nil {
		logrus.Errorf("cant check user: %s", err.Error())
		return err
	}
	if !ok {
		logrus.Infof("cant approve revenue becouse order with id:%d was not created", order.OrderID)
		return fmt.Errorf("order with id:%d was not created", order.OrderID)
	} else {
		err := uc.ordersMatch(order)
		if err != nil {
			logrus.Info(err)
			return err
		}
		if err := uc.trans.RevenueApproval(order); err != nil {
			logrus.Errorf("cant approve revenue: %s", err.Error())
			return err
		}
		logrus.Infof("revenue from order with id:%d was sucsessfully approved", order.OrderID)

		return nil
	}
}
func (uc *UseCase) RevenueDeny(order entities.Order) error {
	ok, err := uc.trans.CheckOrder(order)
	if err != nil {
		logrus.Errorf("cant check order: %s", err.Error())
		return err
	}
	if !ok {
		logrus.Infof("cant deny revenue becouse order with id:%d was not created", order.OrderID)
		return fmt.Errorf("order with id:%d was not created", order.OrderID)
	} else {
		err := uc.ordersMatch(order)
		if err != nil {
			return err
		}
		if err := uc.trans.RevenueDeny(order); err != nil {
			logrus.Errorf("cant deny revenue: %s", err.Error())
			return err
		}
		logrus.Infof("revenue from order with id:%d was sucsessfully denyed", order.OrderID)
		return nil
	}
}

func (uc *UseCase) ordersMatch(order entities.Order) error {
	orderInDB, err := uc.trans.GetOrder(order)
	if err != nil {
		logrus.Errorf("cant match orders: %s", err.Error())
		return err
	}
	if orderInDB != order {
		return fmt.Errorf("order in DB: orderID:%d userID:%d productID:%d price:%f | incoming order: orderID:%d userID:%d productID:%d price:%f", orderInDB.OrderID, orderInDB.UserID, orderInDB.ProductID, orderInDB.Price, order.OrderID, order.UserID, order.ProductID, order.Price)
	}
	return nil
}
