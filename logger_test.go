package logger

import (
	//using the original logrus test implementation, to test the local logger's format
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

//TODO add some more relevant tests
func TestLoggerInit(t *testing.T) {
	hook := NewTestHook("test_service")

	assert.Nil(t, hook.LastEntry())
	assert.Equal(t, 0, len(hook.Entries))

	Infof("[Startup] annotations-monitoring-service is starting")

	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Contains(t, hook.LastEntry().Message, "[Startup] annotations-monitoring-service is starting")
	assert.Equal(t, 1, len(hook.Entries))
}

func TestLogEntryWithTime(t *testing.T) {
	hook := NewTestHook("test_service")

	myExpectedTime := time.Unix(rand.Int63n(time.Now().Unix()), rand.Int63n(1000000000))
	NewEntry("tid_test").WithTime(myExpectedTime).Info("This is a custom time for my event")

	assert.Len(t, hook.Entries, 1)
	assert.Len(t, hook.LastEntry().Data, 2)
	assert.Equal(t, logrus.InfoLevel, hook.LastEntry().Level)
	assert.Equal(t, "This is a custom time for my event", hook.LastEntry().Message)

	myActualTime, err := time.Parse(timestampFormat, hook.LastEntry().Data[fieldKeyTime].(string))
	assert.NoError(t, err)
	assert.WithinDuration(t, myExpectedTime, myActualTime, 0)
	assert.Equal(t, "tid_test", hook.LastEntry().Data["transaction_id"])
}
