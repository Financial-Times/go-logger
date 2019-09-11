package logger

import (
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// The set of these unit tests aim to test the common methods (for both the logger and the entry) but called
// on the logger object. That's why we are looking at the first calls in one call chain.

func TestUPPLoggerWithUUID(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)

	ulog.WithUUID("test-uuid").Info("test info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 1)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "test info message", hook.LastEntry().Message)
	assert.Equal(t, "test-uuid", hook.LastEntry().Data[DefaultKeyUUID])
}

func TestUPPLoggerWithIsValid(t *testing.T) {
	conf := KeyNamesConfig{KeyIsValid: "test-is-valid-key"}
	ulog := NewUPPInfoLogger("test_service", conf)
	hook := test.NewLocal(ulog.Logger)

	ulog.WithValidFlag(true).Info("test info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 1)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "test info message", hook.LastEntry().Message)
	assert.Equal(t, "true", hook.LastEntry().Data[conf.KeyIsValid])
}

func TestUPPLoggerWithTime(t *testing.T) {
	conf := KeyNamesConfig{KeyTime: "test-time-key"}
	ulog := NewUPPInfoLogger("test_service", conf)
	hook := test.NewLocal(ulog.Logger)

	myTime := time.Now()
	ulog.WithTime(myTime).Info("test info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 1)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "test info message", hook.LastEntry().Message)
	assert.Equal(t, myTime.Format(timestampFormat), hook.LastEntry().Data[conf.KeyTime])
}

func TestWithMonitoringEvent(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)

	ulog.WithMonitoringEvent("an-event", "tid_test", "some-content").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 4)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data[DefaultKeyTransactionID])
}

func TestWithMonitoringEventWithConf(t *testing.T) {
	conf := KeyNamesConfig{
		KeyTransactionID:   "test-transaction-id-key",
		KeyEventName:       "test-event-name-key",
		KeyMonitoringEvent: "test-monitoring-key",
		KeyContentType:     "test-content-type-key",
	}
	ulog := NewUPPInfoLogger("test_service", conf)
	hook := test.NewLocal(ulog.Logger)

	ulog.WithMonitoringEvent("an-event", "tid_test", "some-content").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 4)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data[conf.KeyTransactionID])
	assert.Equal(t, "an-event", hook.LastEntry().Data[conf.KeyEventName])
	assert.Equal(t, "true", hook.LastEntry().Data[conf.KeyMonitoringEvent])
	assert.Equal(t, "some-content", hook.LastEntry().Data[conf.KeyContentType])
}

func TestWithTransactionID(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)

	ulog.WithTransactionID("tid_test").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 1)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data[DefaultKeyTransactionID])
}

func TestWithCategorisedEvent(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)

	ulog.WithCategorisedEvent("an-event", "an-event-category", "an-event-msg", "tid_test").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 4)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data[DefaultKeyTransactionID])
	assert.Equal(t, "an-event", hook.LastEntry().Data[DefaultKeyEventName])
	assert.Equal(t, "an-event-category", hook.LastEntry().Data[DefaultKeyEventCategory])
	assert.Equal(t, "an-event-msg", hook.LastEntry().Data[DefaultKeyEventMsg])
}
