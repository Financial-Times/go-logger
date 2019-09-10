package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDefaultKeyNamesConfig(t *testing.T) {
	conf := GetDefaultKeyNamesConfig()
	assert.Equal(t, conf.KeyLogLevel, DefaultKeyLogLevel)
	assert.Equal(t, conf.KeyMsg, DefaultKeyMsg)
	assert.Equal(t, conf.KeyError, DefaultKeyError)
	assert.Equal(t, conf.KeyTime, DefaultKeyTime)
	assert.Equal(t, conf.KeyServiceName, DefaultKeyServiceName)
	assert.Equal(t, conf.KeyTransactionID, DefaultKeyTransactionID)
	assert.Equal(t, conf.KeyUUID, DefaultKeyUUID)
	assert.Equal(t, conf.KeyIsValid, DefaultKeyIsValid)
	assert.Equal(t, conf.KeyEventName, DefaultKeyEventName)
	assert.Equal(t, conf.KeyMonitoringEvent, DefaultKeyMonitoringEvent)
	assert.Equal(t, conf.KeyContentType, DefaultKeyContentType)
	assert.Equal(t, conf.KeyEventCategory, DefaultKeyEventCategory)
	assert.Equal(t, conf.KeyEventMsg, DefaultKeyEventMsg)
}

func TestGetFullKeyNameConfig(t *testing.T) {
	input := KeyNamesConfig{
		KeyLogLevel:    "TEST-LOG-LEVEL",
		KeyMsg:         "test-msg",
		KeyError:       "test-err",
		KeyTime:        "TEST_TIME",
		KeyServiceName: "TestServiceName",
	}
	conf := GetFullKeyNameConfig(input)
	assert.Equal(t, conf.KeyLogLevel, "TEST-LOG-LEVEL")
	assert.Equal(t, conf.KeyMsg, "test-msg")
	assert.Equal(t, conf.KeyError, "test-err")
	assert.Equal(t, conf.KeyTime, "TEST_TIME")
	assert.Equal(t, conf.KeyServiceName, "TestServiceName")
	assert.Equal(t, conf.KeyTransactionID, DefaultKeyTransactionID)
	assert.Equal(t, conf.KeyUUID, DefaultKeyUUID)
	assert.Equal(t, conf.KeyIsValid, DefaultKeyIsValid)
	assert.Equal(t, conf.KeyEventName, DefaultKeyEventName)
	assert.Equal(t, conf.KeyMonitoringEvent, DefaultKeyMonitoringEvent)
	assert.Equal(t, conf.KeyContentType, DefaultKeyContentType)
	assert.Equal(t, conf.KeyEventCategory, DefaultKeyEventCategory)
	assert.Equal(t, conf.KeyEventMsg, DefaultKeyEventMsg)
}
