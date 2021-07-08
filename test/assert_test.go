package test

/*
import (
	"math/rand"
	"testing"
	"time"

	logger "github.com/Financial-Times/go-logger/v2"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAssertHasField(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithField("foo", "bar").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasField("foo", "bar")
	assert.False(t, mockT.Failed())
}

func TestAssertHasFieldFailed(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithField("foo", "bar").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasField("bar", "foo")
	assert.True(t, mockT.Failed())
}

func TestAssertHasError(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithError(assert.AnError).Error()
	e := hook.LastEntry()
	Assert(mockT, e).HasError(assert.AnError)
	assert.False(t, mockT.Failed())
}

func TestAssertHasErrorFailed(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.Error()
	e := hook.LastEntry()
	Assert(mockT, e).HasError(assert.AnError)
	assert.True(t, mockT.Failed())
}

func TestAssertHasFields(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	fields := map[string]interface{}{"foo1": "bar1", "foo2": "bar2"}
	ulog.WithFields(fields).Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasFields(fields)
	assert.False(t, mockT.Failed())
}

func TestAssertHasFieldsFailed(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	fields := map[string]interface{}{"foo1": "bar1", "foo2": "bar2"}
	ulog.WithFields(fields).Info()
	e := hook.LastEntry()
	unexpectedFields := map[string]interface{}{"foo1": "bar1", "foo3": "bar3"}
	Assert(mockT, e).HasFields(unexpectedFields)
	assert.True(t, mockT.Failed())
}

func TestAssertHasMonitoringEvent(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithMonitoringEvent("anEvent", "tid_test", "aContentType").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasMonitoringEvent("anEvent", "tid_test", "aContentType")
	assert.False(t, mockT.Failed())
}

func TestAssertHasMonitoringEventFailedByEvent(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithMonitoringEvent("anEvent", "tid_test", "aContentType").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasMonitoringEvent("anotherEvent", "tid_test", "aContentType")
	assert.True(t, mockT.Failed())
}

func TestAssertHasMonitoringEventFailedByTransactionID(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithMonitoringEvent("anEvent", "tid_test", "aContentType").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasMonitoringEvent("anEvent", "tid_another", "aContentType")
	assert.True(t, mockT.Failed())
}

func TestAssertHasMonitoringEventFailedByContentType(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithMonitoringEvent("anEvent", "tid_test", "aContentType").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasMonitoringEvent("anEvent", "tid_test", "anotherContentType")
	assert.True(t, mockT.Failed())
}

func TestAssertHasTime(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	expectedTime := time.Unix(rand.Int63n(time.Now().Unix()), rand.Int63n(1000000000))
	ulog.WithTransactionID("tid_test").WithTime(expectedTime).Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasTime(expectedTime)
	assert.False(t, mockT.Failed())
}

func TestAssertHasTimeFailed(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	unexpectedTime := time.Date(2003, time.August, 23, 12, 4, 5, 123, time.UTC)
	ulog.WithTransactionID("tid_test").WithTime(time.Now()).Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasTime(unexpectedTime)
	assert.True(t, mockT.Failed())
}

func TestAssertHasTransactionID(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithTransactionID("tid_test").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasTransactionID("tid_test")
	assert.False(t, mockT.Failed())
}

func TestAssertHasTransactionIDFailed(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithTransactionID("tid_test").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasTransactionID("tid_test2")
	assert.True(t, mockT.Failed())
}

func TestAssertHasUUID(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithTransactionID("tid_test").WithUUID("dbc40c07-63ef-4ea3-82d6-a5a5d8747363").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasUUID("dbc40c07-63ef-4ea3-82d6-a5a5d8747363")
	assert.False(t, mockT.Failed())
}

func TestAssertHasUUIDFailed(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithTransactionID("tid_test").WithUUID("dbc40c07-63ef-4ea3-82d6-a5a5d8747363").Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasUUID("c4de1f2c-a7a9-4617-add8-4faacf2fae3e")
	assert.True(t, mockT.Failed())
}

func TestAssertHasValidFlagTrue(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithTransactionID("tid_test").WithValidFlag(true).Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasValidFlag(true)
	assert.False(t, mockT.Failed())
}

func TestAssertHasValidFlagTrueFailed(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithTransactionID("tid_test").WithValidFlag(true).Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasValidFlag(false)
	assert.True(t, mockT.Failed())
}

func TestAssertHasValidFlagFalse(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithTransactionID("tid_test").WithValidFlag(false).Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasValidFlag(false)
	assert.False(t, mockT.Failed())
}

func TestAssertHasValidFlagFalseFailed(t *testing.T) {
	mockT := new(testing.T)
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	ulog.WithTransactionID("tid_test").WithValidFlag(false).Info()
	e := hook.LastEntry()
	Assert(mockT, e).HasValidFlag(true)
	assert.True(t, mockT.Failed())
}

func TestNoDataRace(t *testing.T) {
	ulog := logger.NewUPPInfoLogger("test_service")
	hook := test.NewLocal(ulog.Logger)
	go func() {
		ulog.Info("Something info")
	}()
	time.Sleep(100 * time.Millisecond)

	require.NotNil(t, hook.LastEntry())
	assert.Equal(t, "info", hook.LastEntry().Level.String())
}
*/
