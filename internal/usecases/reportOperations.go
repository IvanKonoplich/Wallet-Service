package usecases

import (
	"errors"
	"github.com/sirupsen/logrus"
)

func (uc *UseCase) GetMonthReport(month string) (string, error) {
	switch month {
	case "January":
	case "February":
	case "March":
	case "April":
	case "May":
	case "June":
	case "July":
	case "August":
	case "September":
	case "October":
	case "November":
	case "December":
	default:
		logrus.Errorf("cant get report: the month is specified incorrectly")
		return "", errors.New("the month is specified incorrectly")
	}
	result, err := uc.trans.GetMonthReport(month)
	if err != nil {
		logrus.Errorf("cant make month report:%s", err.Error())
		return result, err
	}
	logrus.Infof("month report made sucsessfully: %s", result)
	return result, err
}
