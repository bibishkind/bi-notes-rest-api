package main

import (
	"context"
	"github.com/bibishkin/bi-notes-rest-api/pkg/repository/postgres"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logger := getLogger()

	if err := initConfig(); err != nil {
		logger.Fatalf("error intializing config: %s", err.Error())
	}
	logger.Infof("config initialized successfully")

	if err := godotenv.Load(); err != nil {
		logger.Fatalf("error loading env variables: %s", err.Error())
	}
	logger.Infof("env variables successfully loaded")

	postgresCfg := postgres.Config{
		Username:     viper.GetString("db.username"),
		Password:     os.Getenv("DB_PASSWORD"),
		Host:         viper.GetString("db.host"),
		Port:         viper.GetString("db.port"),
		DBName:       viper.GetString("db.name"),
		PoolMaxConns: 10,
	}

	postgresPool, err := postgres.NewConnectionPool(context.Background(), postgresCfg)
	if err != nil {
		logger.Fatalf("database connection error: %s", err.Error())
	}
	logger.Infof("database connection successfully made")

	_ = postgres.NewRepository(postgresPool)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}

func getLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(new(logrus.JSONFormatter))
	logger.SetReportCaller(true)

	return logger
}
