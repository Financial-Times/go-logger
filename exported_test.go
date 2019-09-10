package logger

import (
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestWithMonitoringEvent(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)

	ulog.WithMonitoringEvent("an-event", "tid_test", "some-content").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 4)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}

func TestWithTransactionID(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)

	ulog.WithTransactionID("tid_test").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 1)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}

func TestWithCategorisedEvent(t *testing.T) {
	ulog := NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)

	ulog.WithCategorisedEvent("an-event", "an-event-category", "an-event-msg", "tid_test").Info("a info message")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 4)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "a info message", hook.LastEntry().Message)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
	assert.Equal(t, "an-event", hook.LastEntry().Data["event"])
	assert.Equal(t, "an-event-category", hook.LastEntry().Data["event_category"])
	assert.Equal(t, "an-event-msg", hook.LastEntry().Data["event_msg"])
}
