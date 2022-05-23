package core

import (
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func LoadConfig() {
	if err := godotenv.Load(); err != nil {
		logrus.Warnf("Error loading .env file: %v", err)
	}
	logrus.Info("Loaded .env file")
}
