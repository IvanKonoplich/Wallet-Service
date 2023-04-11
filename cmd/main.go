package main

import (
	"avitoTest/internal/app"
	"avitoTest/internal/infrastructure/repository/storage"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatalf("error while reading config file:%s", err.Error())
	}
	postgresPassword := initGoDotEnv()
	postgresConfig := storage.ConfigDB{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.DBName"),
		Password: postgresPassword,
		SSLMode:  viper.GetString("db.SSLMode"),
	}
	app.Run(postgresConfig)
}

func initConfig() error {
	viper.SetDefault("port", "8000")
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("config")
	return viper.ReadInConfig()
}

func initGoDotEnv() string {
	err := godotenv.Load()
	if err != nil {
		logrus.Fatal("Error loading .env file")
	}

	return os.Getenv("POSTGRES_PASSWORD")
}
