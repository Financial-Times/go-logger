package logger

import (
	"encoding/json"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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
	f := newFTJSONFormatter(testServiceName, GetDefaultKeyNamesConfig())
	ulog := NewUnstructuredLogger()
	e := ulog.WithMonitoringEvent(testEvent, testTID, testContentType).WithError(errors.New(testErrMsg))
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e)
	assert.NoError(t, err)

	var logLine map[string]string
	err = json.Unmarshal(logLineBytes, &logLine)
	assert.NoError(t, err)
	assert.Len(t, logLine, 9)

	actualTime, err := time.Parse(timestampFormat, logLine[DefaultKeyTime])
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), actualTime, 2*time.Second)

	assert.Equal(t, testServiceName, logLine[DefaultKeyServiceName])
	assert.Equal(t, testEvent, logLine[DefaultKeyEventName])
	assert.Equal(t, testTID, logLine[DefaultKeyTransactionID])
	assert.Equal(t, testContentType, logLine[DefaultKeyContentType])
	assert.Equal(t, testErrMsg, logLine[DefaultKeyError])
	assert.Equal(t, testMsg, logLine[logrus.FieldKeyMsg])
	assert.Equal(t, logrus.InfoLevel.String(), logLine[logrus.FieldKeyLevel])
	assert.Equal(t, "true", logLine[DefaultKeyMonitoringEvent])
}

func TestFtJSONFormatterWithConf(t *testing.T) {
	conf := KeyNamesConfig{
		KeyLogLevel: "test-log-level-key",
		KeyMsg:      "test-msg-key",
		//KeyError:           "test-err-key",
		KeyTime:            "test-time-key",
		KeyServiceName:     "test-service-name",
		KeyTransactionID:   "test-trans-id",
		KeyEventName:       "test-event-name-key",
		KeyMonitoringEvent: "test-monitoring-event-key",
		KeyContentType:     "test-content-type",
	}
	f := newFTJSONFormatter(testServiceName, GetFullKeyNameConfig(conf))
	ulog := NewUPPInfoLogger(testServiceName, conf)
	e := ulog.WithMonitoringEvent(testEvent, testTID, testContentType).WithError(errors.New(testErrMsg))
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e)
	assert.NoError(t, err)

	var logLine map[string]string
	err = json.Unmarshal(logLineBytes, &logLine)
	assert.NoError(t, err)
	assert.Len(t, logLine, 9)

	actualTime, err := time.Parse(timestampFormat, logLine[conf.KeyTime])
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), actualTime, 2*time.Second)

	assert.Equal(t, testServiceName, logLine[conf.KeyServiceName])
	assert.Equal(t, testEvent, logLine[conf.KeyEventName])
	assert.Equal(t, testTID, logLine[conf.KeyTransactionID])
	assert.Equal(t, testContentType, logLine[conf.KeyContentType])
	assert.Equal(t, testErrMsg, logLine["error"])
	assert.Equal(t, testMsg, logLine[conf.KeyMsg])
	assert.Equal(t, logrus.InfoLevel.String(), logLine[conf.KeyLogLevel])
	assert.Equal(t, "true", logLine[conf.KeyMonitoringEvent])
}

func TestFtJSONFormatterWithLogTimeField(t *testing.T) {
	myExpectedTime := time.Unix(rand.Int63n(time.Now().Unix()), rand.Int63n(1000000000))
	f := newFTJSONFormatter(testServiceName, GetDefaultKeyNamesConfig())
	ulog := NewUnstructuredLogger()
	e := ulog.WithMonitoringEvent(testEvent, testTID, testContentType).WithTime(myExpectedTime).
		WithError(errors.New(testErrMsg))
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e)
	assert.NoError(t, err)

	var logLine map[string]string
	err = json.Unmarshal(logLineBytes, &logLine)
	assert.NoError(t, err)
	assert.Len(t, logLine, 9)

	myActualTime, err := time.Parse(timestampFormat, logLine[DefaultKeyTime])
	assert.NoError(t, err)
	assert.WithinDuration(t, myExpectedTime, myActualTime, 0)

	assert.Equal(t, testServiceName, logLine[DefaultKeyServiceName])
	assert.Equal(t, testEvent, logLine[DefaultKeyEventName])
	assert.Equal(t, testTID, logLine[DefaultKeyTransactionID])
	assert.Equal(t, testContentType, logLine[DefaultKeyContentType])
	assert.Equal(t, testErrMsg, logLine["error"])
	assert.Equal(t, testMsg, logLine[logrus.FieldKeyMsg])
	assert.Equal(t, logrus.InfoLevel.String(), logLine[logrus.FieldKeyLevel])
	assert.Equal(t, "true", logLine[DefaultKeyMonitoringEvent])
}

func TestLoggerWithoutInitialisation(t *testing.T) {
	f := newFTJSONFormatter("", GetDefaultKeyNamesConfig())
	ulog := NewUnstructuredLogger()
	e := ulog.WithMonitoringEvent(testEvent, testTID, testContentType).WithError(errors.New(testErrMsg))
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.ErrorLevel

	logLineBytes, err := f.Format(e)

	assert.Empty(t, logLineBytes)
	assert.EqualError(t, err, "UPP log formatter is not initialised with service name")
}

func TestFtJSONFormatterWithStructuredEvent(t *testing.T) {
	f := newFTJSONFormatter(testServiceName, GetDefaultKeyNamesConfig())
	ulog := NewUnstructuredLogger()
	e := ulog.WithCategorisedEvent(testEvent, "event-category", "event-msg", testTID).
		WithError(errors.New(testErrMsg))
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e)
	assert.NoError(t, err)

	var logLine map[string]string
	err = json.Unmarshal(logLineBytes, &logLine)
	assert.NoError(t, err)
	assert.Len(t, logLine, 9)

	actualTime, err := time.Parse(timestampFormat, logLine[DefaultKeyTime])
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), actualTime, 2*time.Second)

	assert.Equal(t, testServiceName, logLine[DefaultKeyServiceName])
	assert.Equal(t, testEvent, logLine[DefaultKeyEventName])
	assert.Equal(t, testTID, logLine[DefaultKeyTransactionID])
	assert.Equal(t, "event-category", logLine[DefaultKeyEventCategory])
	assert.Equal(t, "event-msg", logLine[DefaultKeyEventMsg])
	assert.Equal(t, testErrMsg, logLine["error"])
	assert.Equal(t, testMsg, logLine[logrus.FieldKeyMsg])
	assert.Equal(t, logrus.InfoLevel.String(), logLine[logrus.FieldKeyLevel])
}
