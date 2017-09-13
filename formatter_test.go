package logger

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	testServiceName = "test-service-api"
	testEvent       = "apocalypse"
	testTID         = "tid_test"
	testContentType = "lionel-barber-biography"
	testErrMsg      = "the world is over"
	testMsg         = "happy ending"
)

func TestFtJSONFormatter(t *testing.T) {
	f := newFTJSONFormatter(testServiceName)
	e := NewMonitoringEntry(testEvent, testTID, testContentType).WithError(errors.New(testErrMsg))
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e)
	assert.NoError(t, err)

	var logLine map[string]string
	err = json.Unmarshal(logLineBytes, &logLine)
	assert.NoError(t, err)
	assert.Len(t, logLine, 9)

	actualTime, err := time.Parse(timestampFormat, logLine[fieldKeyTime])
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), actualTime, 2*time.Second)

	assert.Equal(t, testServiceName, logLine[fieldKeyServiceName])
	assert.Equal(t, testEvent, logLine["event"])
	assert.Equal(t, testTID, logLine["transaction_id"])
	assert.Equal(t, testContentType, logLine["content_type"])
	assert.Equal(t, testErrMsg, logLine["error"])
	assert.Equal(t, testMsg, logLine[logrus.FieldKeyMsg])
	assert.Equal(t, logrus.InfoLevel.String(), logLine[logrus.FieldKeyLevel])
	assert.Equal(t, "true", logLine["monitoring_event"])
}
