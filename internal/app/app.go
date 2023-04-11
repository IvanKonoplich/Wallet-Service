package app

import (
	"avitoTest/internal/controllers"
	"avitoTest/internal/infrastructure/repository/storage"
	"avitoTest/internal/infrastructure/repository/transactionsMaker"
	"avitoTest/internal/usecases"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Run(postgresConfig storage.ConfigDB) {
	postgresDB, err := storage.OpenDBConnection(postgresConfig)
	if err != nil {
		logrus.Fatalf("error opening postgres connection:%s", err.Error())
	}
	repos := storage.NewStorage(postgresDB)
	txMaker := transactionsMaker.NewTransactionMaker(repos)
	useCase := usecases.New(txMaker)
	controller := controllers.New(useCase)
	router := controller.InitRouter()
	server := new(controllers.Server)
	if err := server.RunServer(viper.GetString("port"), router); err != nil {
		logrus.Fatalf("error while starting server:%s", err.Error())
	}
	logrus.Info("starting server...")
}
