package core

import (
	"time"

	"github.com/sirupsen/logrus"
)

func InitLogger() {
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		ForceQuote:      true,
		FullTimestamp:   true,
		TimestampFormat: time.Kitchen,
		PadLevelText:    true,
	})

	logrus.Info("Logger initialized")
}
