package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoggerInit(t *testing.T) {
	hook := NewTestHook("test_service")

	assert.Nil(t, hook.LastEntry())
	assert.Equal(t, 0, len(hook.Entries))

	Infof("[Startup] annotations-monitoring-service is starting")

	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Contains(t, hook.LastEntry().Message, "[Startup] annotations-monitoring-service is starting")
	assert.Equal(t, 1, len(hook.Entries))
}

func TestInitLoggerFatal(t *testing.T) {
	hook := test.NewLocal(log)

	InitLogger("test_service", "a-level-that-do-not-exist")

	assert.Equal(t, logrus.ErrorLevel, hook.LastEntry().Level)
	assert.Contains(t, hook.LastEntry().Message, "Incorrect log level. Using INFO instead.")
	assert.Equal(t, 1, len(hook.Entries))
	assert.Equal(t, "a-level-that-do-not-exist", hook.LastEntry().Data["logLevel"])
	assert.EqualError(t, hook.LastEntry().Data["error"].(error), `not a valid logrus Level: "a-level-that-do-not-exist"`)
}
