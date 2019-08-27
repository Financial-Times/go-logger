package test

import (
	"github.com/Financial-Times/go-logger/v2"
	"github.com/sirupsen/logrus/hooks/test"
)

func NewTestHook(serviceName string) *test.Hook {
	logger.InitDefaultLogger(serviceName)
	return test.NewLocal(logger.Logger())
}
