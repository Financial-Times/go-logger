package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

func TestLogEntryWithTime(t *testing.T) {
	hook := NewTestHook("test_service")

	myExpectedTime := time.Unix(rand.Int63n(time.Now().Unix()), rand.Int63n(1000000000))
	WithMonitoringEvent("an-event", "tid_test", "some-content").WithTime(myExpectedTime).Info("This is a custom time for my event")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "This is a custom time for my event", hook.LastEntry().Message)

	myActualTime, err := time.Parse(timestampFormat, hook.LastEntry().Data[fieldKeyTime].(string))
	assert.NoError(t, err)
	assert.WithinDuration(t, myExpectedTime, myActualTime, 0)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}

func TestLogEntryWithTransactionID(t *testing.T) {
	hook := NewTestHook("test_service")

	expectedUUID := "50484f2a-a51d-42d8-8deb-11a1d25e6b45"

	WithMonitoringEvent("an-event", "tid_test", "some-content").WithUUID(expectedUUID).Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, expectedUUID, hook.LastEntry().Data["uuid"])
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}

func TestLogEntryWithValidFlagTrue(t *testing.T) {
	hook := NewTestHook("test_service")

	WithMonitoringEvent("an-event", "tid_test", "some-content").WithValidFlag(true).Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	actualValidFlag, err := strconv.ParseBool(hook.LastEntry().Data["isValid"].(string))
	assert.NoError(t, err)
	assert.True(t, actualValidFlag)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}

func TestLogEntryWithValidFlagFalse(t *testing.T) {
	hook := NewTestHook("test_service")

	WithMonitoringEvent("an-event", "tid_test", "some-content").WithValidFlag(false).Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 5)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	actualValidFlag, err := strconv.ParseBool(hook.LastEntry().Data["isValid"].(string))
	assert.NoError(t, err)
	assert.False(t, actualValidFlag)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}