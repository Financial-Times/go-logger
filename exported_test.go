package logger

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithMonitoringEvent(t *testing.T) {
	hook := NewTestHook("test_service")

	WithMonitoringEvent("an-event", "tid_test", "some-content").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 4)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}

func TestWithTransactionID(t *testing.T) {
	hook := NewTestHook("test_service")

	WithTransactionID("tid_test").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 1)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}
