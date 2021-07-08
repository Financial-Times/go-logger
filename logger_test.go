package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestNewUPPLogger(t *testing.T) {
	ilog := NewUPPLogger("test_service", "info")
	ulog, _ := ilog.(*UPPLogger)
	hook := test.NewLocal(ulog.Logger)

	assert.Nil(t, hook.LastEntry())
	assert.Equal(t, 0, len(hook.Entries))

	ulog.Infof("[Startup] annotations-monitoring-service is starting")

	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Contains(t, hook.LastEntry().Message, "[Startup] annotations-monitoring-service is starting")
	assert.Equal(t, 1, len(hook.Entries))
}

func TestNewUPPLoggerDefaultConf(t *testing.T) {
	ilog := NewUPPInfoLogger("test_service")
	ulog, _ := ilog.(*UPPLogger)
	assert.Equal(t, ulog.keyConf, GetDefaultKeyNamesConfig())
}

func TestNewUPPLoggerWithConf(t *testing.T) {
	conf := KeyNamesConfig{
		KeyLogLevel: "TEST_LOG_LEVEL",
		KeyMsg:      "TEST_LOG_MSG",
	}
	ilog := NewUPPLogger("test_service", "info", conf)
	ulog, _ := ilog.(*UPPLogger)
	assert.Equal(t, "TEST_LOG_LEVEL", ulog.keyConf.KeyLogLevel)
	assert.Equal(t, "TEST_LOG_MSG", ulog.keyConf.KeyMsg)
	assert.Equal(t, DefaultKeyTransactionID, ulog.keyConf.KeyTransactionID)
}

func TestNewUPPInfoLogger(t *testing.T) {
	ilog := NewUPPInfoLogger("test_service")
	ulog, _ := ilog.(*UPPLogger)
	assert.Equal(t, logrus.InfoLevel, ulog.Logger.Level)
	assert.Equal(t, ulog.keyConf, GetDefaultKeyNamesConfig())
}

func TestNewUPPInfoLoggerWithConf(t *testing.T) {
	conf := KeyNamesConfig{
		KeyLogLevel: "TEST_LOG_LEVEL",
		KeyMsg:      "TEST_LOG_MSG",
	}
	ilog := NewUPPInfoLogger("test_service", conf)
	ulog, _ := ilog.(*UPPLogger)
	assert.Equal(t, logrus.InfoLevel, ulog.Logger.Level)
	assert.Equal(t, "TEST_LOG_LEVEL", ulog.keyConf.KeyLogLevel)
	assert.Equal(t, "TEST_LOG_MSG", ulog.keyConf.KeyMsg)
	assert.Equal(t, DefaultKeyTransactionID, ulog.keyConf.KeyTransactionID)
}

func TestNewUnstructuredLogger(t *testing.T) {
	ilog := NewUnstructuredLogger()
	ulog, _ := ilog.(*UPPLogger)
	assert.Equal(t, ulog.keyConf, GetDefaultKeyNamesConfig())
}
