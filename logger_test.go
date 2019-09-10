package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestNewUPPLogger(t *testing.T) {
	ulog := NewUPPLogger("test_service", "info")
	hook := test.NewLocal(ulog.Logger)

	assert.Nil(t, hook.LastEntry())
	assert.Equal(t, 0, len(hook.Entries))

	ulog.Infof("[Startup] annotations-monitoring-service is starting")

	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Contains(t, hook.LastEntry().Message, "[Startup] annotations-monitoring-service is starting")
	assert.Equal(t, 1, len(hook.Entries))
}

func TestNewUPPLoggerDefaultConf(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	assert.Equal(t, ulog.keyConf, GetDefaultKeyNamesConfig())
}

func TestNewUPPLoggerWithConf(t *testing.T) {
	conf := KeyNamesConfig{
		KeyLogLevel: "TEST_LOG_LEVEL",
		KeyMsg:      "TEST_LOG_MSG",
	}
	ulog := NewUPPLogger("test_service", "info", conf)
	assert.Equal(t, "TEST_LOG_LEVEL", ulog.keyConf.KeyLogLevel)
	assert.Equal(t, "TEST_LOG_MSG", ulog.keyConf.KeyMsg)
	assert.Equal(t, DefaultKeyTransactionID, ulog.keyConf.KeyTransactionID)
}

func TestNewUPPInfoLogger(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	assert.Equal(t, logrus.InfoLevel, ulog.Logger.Level)
	assert.Equal(t, ulog.keyConf, GetDefaultKeyNamesConfig())
}

func TestNewUPPInfoLoggerWithConf(t *testing.T) {
	conf := KeyNamesConfig{
		KeyLogLevel: "TEST_LOG_LEVEL",
		KeyMsg:      "TEST_LOG_MSG",
	}
	ulog := NewUPPInfoLogger("test_service", conf)
	assert.Equal(t, logrus.InfoLevel, ulog.Logger.Level)
	assert.Equal(t, "TEST_LOG_LEVEL", ulog.keyConf.KeyLogLevel)
	assert.Equal(t, "TEST_LOG_MSG", ulog.keyConf.KeyMsg)
	assert.Equal(t, DefaultKeyTransactionID, ulog.keyConf.KeyTransactionID)
}

func TestNewUnstructuredLogger(t *testing.T) {
	ulog := NewUnstructuredLogger()
	assert.Equal(t, ulog.keyConf, GetDefaultKeyNamesConfig())
}
