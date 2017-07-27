package logger

import (
	"github.com/Sirupsen/logrus/hooks/test"
)

func NewTestHook(serviceName string) *test.Hook {
	InitDefaultLogger(serviceName)
	return test.NewLocal(logger.log)
}
