package helpers

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func InitializeLogging() *logrus.Logger {
	l := logrus.New()

	l.SetFormatter(&logrus.JSONFormatter{
		DataKey:     uuid.NewString(),
		PrettyPrint: false,
	})

	l.SetLevel(logrus.DebugLevel)

	return l
}
