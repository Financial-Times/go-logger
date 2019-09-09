package test

import (
	"strconv"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

// LoggingAssert struct exposes convenient assert methods for UPP logger specific log entries.
type LoggingAssert struct {
	t     *testing.T
	entry *logrus.Entry
}

func Assert(t *testing.T, entry *logrus.Entry) *LoggingAssert {
	return &LoggingAssert{t, entry}
}

func (a *LoggingAssert) HasField(key string, value interface{}) *LoggingAssert {
	assert.Equal(a.t, value, a.entry.Data[key])
	return a
}

func (a *LoggingAssert) HasFields(fields map[string]interface{}) *LoggingAssert {
	for k, v := range fields {
		a.HasField(k, v)
	}
	return a
}

func (a *LoggingAssert) HasMonitoringEvent(expectedEventName, expectedTID, expectedContentType string) *LoggingAssert {
	return a.HasField("event", expectedEventName).
		HasTransactionID(expectedTID).
		HasField("content_type", expectedContentType).
		HasField("monitoring_event", "true")
}

func (a *LoggingAssert) HasValidFlag(expectedFlag bool) *LoggingAssert {
	return a.HasField("isValid", strconv.FormatBool(expectedFlag))
}

func (a *LoggingAssert) HasTransactionID(expectedTID string) *LoggingAssert {
	return a.HasField("transaction_id", expectedTID)
}

func (a *LoggingAssert) HasUUID(expectedUUID string) *LoggingAssert {
	return a.HasField("uuid", expectedUUID)
}

func (a *LoggingAssert) HasTime(expectedTime time.Time) *LoggingAssert {
	return a.HasField("@time", expectedTime.Format(time.RFC3339Nano))
}

func (a *LoggingAssert) HasError(expectedErr error) *LoggingAssert {
	return a.HasField(logrus.ErrorKey, expectedErr)
}
