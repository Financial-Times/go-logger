package logger

import (
	"github.com/sirupsen/logrus/hooks/test"
)

func NewTestHook(serviceName string) *test.Hook {
	InitDefaultLogger(serviceName)
	return test.NewLocal(log)
}
