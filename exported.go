package logger

import (
	"strconv"
	"time"
)

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (ulog *UPPLogger) WithField(key string, value interface{}) *LogEntry {
	return &LogEntry{ulog, ulog.Logger.WithField(key, value)}
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (ulog *UPPLogger) WithFields(fields map[string]interface{}) *LogEntry {
	return &LogEntry{ulog, ulog.Logger.WithFields(fields)}
}

// WithTransactionID creates an entry from the standard logger and adds transaction_id field to it.
func (ulog *UPPLogger) WithTransactionID(tid string) *LogEntry {
	return ulog.WithField(ulog.keyConf.KeyTransactionID, tid)
}

// WithError creates an entry from the standard logger and adds an error field to it.
func (ulog *UPPLogger) WithError(err error) *LogEntry {
	return ulog.WithField(ulog.keyConf.KeyError, err)
}

// WithUUID creates an entry from the standard logger and adds an uuid field to it.
func (ulog *UPPLogger) WithUUID(uuid string) *LogEntry {
	return ulog.WithField(ulog.keyConf.KeyUUID, uuid)
}

// WithValidFlag creates an entry from the standard logger and adds an "is valid" field to it.
func (ulog *UPPLogger) WithValidFlag(isValid bool) *LogEntry {
	return ulog.WithField(ulog.keyConf.KeyIsValid, strconv.FormatBool(isValid))
}

// WithTime creates an entry from the standard logger and adds an time field to it.
func (ulog *UPPLogger) WithTime(time time.Time) *LogEntry {
	return ulog.WithField(ulog.keyConf.KeyTime, time.Format(timestampFormat))
}

// WithMonitoringEvent creates an entry from the standard logger and adds monitoring event fields to it.
// The monitoring event fields are "monitoring_event", "event" and "content_type".
func (ulog *UPPLogger) WithMonitoringEvent(eventName, tid, contentType string) *LogEntry {
	e := ulog.WithFields(
		map[string]interface{}{
			ulog.keyConf.KeyMonitoringEvent: "true",
			ulog.keyConf.KeyEventName:       eventName,
			ulog.keyConf.KeyContentType:     contentType,
		})
	return e.WithTransactionID(tid)
}

// WithMonitoringEvent creates an entry from the standard logger and adds categorised event fields to it.
// The added fields are "monitoring_event", "event" and "content_type".
func (ulog *UPPLogger) WithCategorisedEvent(eventName, eventCategory, eventMsg, tid string) *LogEntry {
	e := ulog.WithFields(
		map[string]interface{}{
			ulog.keyConf.KeyEventName:     eventName,
			ulog.keyConf.KeyEventCategory: eventCategory,
			ulog.keyConf.KeyEventMsg:      eventMsg,
		})
	return e.WithTransactionID(tid)
}
