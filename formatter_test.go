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
	ilog := NewUnstructuredLogger()
	ulog, _ := ilog.(*UPPLogger)
	ientry := ulog.
		WithMonitoringEvent(testEvent, testTID, testContentType).
		WithError(errors.New(testErrMsg))

	e, _ := ientry.(*UPPLogEntry)
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e.Entry)
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
		KeyLogLevel:        "test-log-level-key",
		KeyMsg:             "test-msg-key",
		KeyError:           "test-err-key",
		KeyTime:            "test-time-key",
		KeyServiceName:     "test-service-name",
		KeyTransactionID:   "test-trans-id",
		KeyEventName:       "test-event-name-key",
		KeyMonitoringEvent: "test-monitoring-event-key",
		KeyContentType:     "test-content-type",
	}
	f := newFTJSONFormatter(testServiceName, GetFullKeyNameConfig(conf))
	ilog := NewUPPInfoLogger(testServiceName, conf)
	ulog, _ := ilog.(*UPPLogger)
	ientry := ulog.
		WithMonitoringEvent(testEvent, testTID, testContentType).
		WithError(errors.New(testErrMsg))
	e, _ := ientry.(*UPPLogEntry)
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e.Entry)
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
	assert.Equal(t, testErrMsg, logLine[conf.KeyError])
	assert.Equal(t, testMsg, logLine[conf.KeyMsg])
	assert.Equal(t, logrus.InfoLevel.String(), logLine[conf.KeyLogLevel])
	assert.Equal(t, "true", logLine[conf.KeyMonitoringEvent])
}

func TestFtJSONFormatterWithLogTimeField(t *testing.T) {
	myExpectedTime := time.Unix(rand.Int63n(time.Now().Unix()), rand.Int63n(1000000000))
	f := newFTJSONFormatter(testServiceName, GetDefaultKeyNamesConfig())
	ilog := NewUnstructuredLogger()
	ulog, _ := ilog.(*UPPLogger)
	ientry := ulog.
		WithMonitoringEvent(testEvent, testTID, testContentType).
		WithTime(myExpectedTime).
		WithError(errors.New(testErrMsg))

	e, _ := ientry.(*UPPLogEntry)
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e.Entry)
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
	ilog := NewUnstructuredLogger()
	ulog, _ := ilog.(*UPPLogger)
	ientry := ulog.WithMonitoringEvent(testEvent, testTID, testContentType).WithError(errors.New(testErrMsg))
	e, _ := ientry.(*UPPLogEntry)
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.ErrorLevel

	logLineBytes, err := f.Format(e.Entry)

	assert.Empty(t, logLineBytes)
	assert.EqualError(t, err, "UPP log formatter is not initialised with service name")
}

func TestFtJSONFormatterWithStructuredEvent(t *testing.T) {
	f := newFTJSONFormatter(testServiceName, GetDefaultKeyNamesConfig())
	ilog := NewUnstructuredLogger()
	ulog, _ := ilog.(*UPPLogger)
	ientry := ulog.
		WithCategorisedEvent(testEvent, "event-category", "event-msg", testTID).
		WithError(errors.New(testErrMsg))
	e, _ := ientry.(*UPPLogEntry)
	e.Time = time.Now()
	e.Message = testMsg
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e.Entry)
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

func TestFtJSONFormatterEmptyVals(t *testing.T) {
	f := newFTJSONFormatter(testServiceName, GetDefaultKeyNamesConfig())
	ilog := NewUnstructuredLogger()
	ulog, _ := ilog.(*UPPLogger)
	fields := map[string]interface{}{
		"key-with-val": "val",
		"key-empty":    "",
		"key-nil":      nil,
	}
	ientry := ulog.WithFields(fields).WithTransactionID(testTID)
	e, _ := ientry.(*UPPLogEntry)
	e.Time = time.Now()
	e.Message = ""
	e.Level = logrus.InfoLevel

	logLineBytes, err := f.Format(e.Entry)
	assert.NoError(t, err)

	var logLine map[string]string
	err = json.Unmarshal(logLineBytes, &logLine)
	assert.NoError(t, err)
	assert.Len(t, logLine, 5)

	actualTime, err := time.Parse(timestampFormat, logLine[DefaultKeyTime])
	assert.NoError(t, err)
	assert.WithinDuration(t, time.Now(), actualTime, 2*time.Second)

	assert.Equal(t, testServiceName, logLine[DefaultKeyServiceName])
	assert.Equal(t, testTID, logLine[DefaultKeyTransactionID])
	assert.Equal(t, "val", logLine["key-with-val"])
	assert.Equal(t, logrus.InfoLevel.String(), logLine[logrus.FieldKeyLevel])

	assert.NotContains(t, logLine, "key-empty")
	assert.NotContains(t, logLine, "key-nil")
	assert.NotContains(t, logLine, logrus.FieldKeyMsg)
}
