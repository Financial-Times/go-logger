package logger

import (
	//using the original logrus test implementation, to test the local logger's format
	"github.com/Sirupsen/logrus"
	testLogger "github.com/Sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

//TODO add some more relevant tests
func TestLoggerInit(t *testing.T) {

	InitDefaultLogger("test_service")

	hook := testLogger.NewLocal(logger.Logger)
	assert.Nil(t, hook.LastEntry())
	assert.Equal(t, 0, len(hook.Entries))

	Infof(nil, "[Startup] annotations-monitoring-service is starting")

	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Contains(t, hook.LastEntry().Message, "[Startup] annotations-monitoring-service is starting")
	assert.Equal(t, 1, len(hook.Entries))
}
