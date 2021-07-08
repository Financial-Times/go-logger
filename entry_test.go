package logger

import (
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

// The set of these unit tests aim to test the common methods (for both the logger and the entry) but called
// on the entry object. That's why we are looking at the second calls in one call chain.
func TestLogEntryWithTime(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	log, _ := ulog.(*UPPLogger)
	hook := test.NewLocal(log.Logger)

	myExpectedTime := time.Unix(rand.Int63n(time.Now().Unix()), rand.Int63n(1000000000))
	ulog.WithMonitoringEvent("an-event", "tid_test", "some-content").WithTime(myExpectedTime).Info("This is a custom time for my event")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "This is a custom time for my event", hook.LastEntry().Message)

	myActualTime, err := time.Parse(timestampFormat, hook.LastEntry().Data[DefaultKeyTime].(string))
	assert.NoError(t, err)
	assert.WithinDuration(t, myExpectedTime, myActualTime, 0)
	assert.Equal(t, "tid_test", hook.LastEntry().Data[DefaultKeyTransactionID])
}

func TestLogEntryWithTransactionID(t *testing.T) {
	conf := KeyNamesConfig{KeyTransactionID: "test-trans-id"}
	ilog := NewUPPInfoLogger("test_service", conf)
	ulog, _ := ilog.(*UPPLogger)
	hook := test.NewLocal(ulog.Logger)
	expectedUUID := "50484f2a-a51d-42d8-8deb-11a1d25e6b45"

	ulog.WithMonitoringEvent("an-event", "tid_test", "some-content").WithUUID(expectedUUID).Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, expectedUUID, hook.LastEntry().Data[DefaultKeyUUID])
	assert.Equal(t, "tid_test", hook.LastEntry().Data[conf.KeyTransactionID])
}

func TestLogEntryWithValidFlagTrue(t *testing.T) {
	ilog := NewUPPInfoLogger("test_service")
	ulog, _ := ilog.(*UPPLogger)
	hook := test.NewLocal(ulog.Logger)

	ulog.WithMonitoringEvent("an-event", "tid_test", "some-content").WithValidFlag(true).Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	isValid, err := strconv.ParseBool(hook.LastEntry().Data[DefaultKeyIsValid].(string))
	assert.NoError(t, err)
	assert.True(t, isValid)
	assert.Equal(t, "tid_test", hook.LastEntry().Data[DefaultKeyTransactionID])
}

func TestLogEntryWithValidFlagFalse(t *testing.T) {
	ilog := NewUPPInfoLogger("test_service")
	ulog, _ := ilog.(*UPPLogger)
	hook := test.NewLocal(ulog.Logger)

	ulog.WithMonitoringEvent("an-event", "tid_test", "some-content").
		WithValidFlag(false).Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	isValid, err := strconv.ParseBool(hook.LastEntry().Data[DefaultKeyIsValid].(string))
	assert.NoError(t, err)
	assert.False(t, isValid)
	assert.Equal(t, "tid_test", hook.LastEntry().Data[DefaultKeyTransactionID])
}

func TestLogEntryWithMonitoringEvent(t *testing.T) {
	conf := KeyNamesConfig{KeyEventName: "test-event-name"}
	ilog := NewUPPInfoLogger("test_service", conf)
	ulog, _ := ilog.(*UPPLogger)
	hook := test.NewLocal(ulog.Logger)

	ulog.WithUUID("test-uuid-value").
		WithMonitoringEvent("an-event", "tid_test", "some-content").
		Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "an-event", hook.LastEntry().Data[conf.KeyEventName])
	assert.Equal(t, "true", hook.LastEntry().Data[DefaultKeyMonitoringEvent])
	assert.Equal(t, "some-content", hook.LastEntry().Data[DefaultKeyContentType])
	assert.Equal(t, "tid_test", hook.LastEntry().Data[DefaultKeyTransactionID])
}

func TestLogEntryWithCategorisedEvent(t *testing.T) {
	conf := KeyNamesConfig{KeyEventCategory: "test-event-category-key"}
	ilog := NewUPPInfoLogger("test_service", conf)
	ulog, _ := ilog.(*UPPLogger)
	hook := test.NewLocal(ulog.Logger)

	ulog.WithUUID("test-uuid-value").
		WithCategorisedEvent("test-event", "test-category", "test-event-msg", "test-tid").
		Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "test-event", hook.LastEntry().Data[DefaultKeyEventName])
	assert.Equal(t, "test-category", hook.LastEntry().Data[conf.KeyEventCategory])
	assert.Equal(t, "test-event-msg", hook.LastEntry().Data[DefaultKeyEventMsg])
	assert.Equal(t, "test-tid", hook.LastEntry().Data[DefaultKeyTransactionID])
}
