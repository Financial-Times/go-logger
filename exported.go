package logger

import (
	"github.com/sirupsen/logrus"
)

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func (ulog *UPPLogger) WithError(err error) LogEntry {
	return &logEntry{ulog.Logger.WithField(logrus.ErrorKey, err)}
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (ulog *UPPLogger) WithField(key string, value interface{}) LogEntry {
	return &logEntry{ulog.Logger.WithField(key, value)}
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (ulog *UPPLogger) WithFields(fields map[string]interface{}) LogEntry {
	return &logEntry{ulog.Logger.WithFields(fields)}
}

func (ulog *UPPLogger) WithTransactionID(tid string) LogEntry {
	return &logEntry{ulog.Logger.WithField("transaction_id", tid)}
}

func (ulog *UPPLogger) WithMonitoringEvent(eventName, tid, contentType string) LogEntry {
	e := &logEntry{
		ulog.WithField("monitoring_event", "true").
			WithField("event", eventName).
			WithField("content_type", contentType),
	}
	return e.WithTransactionID(tid)
}
