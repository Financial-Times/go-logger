package test

import (
	"github.com/Financial-Times/go-logger"
	"github.com/sirupsen/logrus/hooks/test"
)

func NewTestHook(serviceName string) *test.Hook {
	logger.InitDefaultLogger(serviceName)
	return test.NewLocal(logger.Logger())
}
