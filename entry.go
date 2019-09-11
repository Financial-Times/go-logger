package logger

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// logEntry is wrapper around logrus.Entry with a few methods adding UPP specific keys to the log entries.
// logEntry is the final or intermediate logging entry. It's finally logged when Debug, Info,
// Warn, Error, Fatal or Panic is called on it.
type LogEntry struct {
	ulog *UPPLogger
	*logrus.Entry
}

// WithField for LogEntry is wrapper around WithField of logrus.Entry
// that let us return our object when chaining the calls to WithField.
func (entry *LogEntry) WithField(key string, value interface{}) *LogEntry {
	return &LogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(key, value)}
}

// WithFields for LogEntry is wrapper around WithFields of logrus.Entry
// that let us return our object when chaining the calls to WithFields.
func (entry *LogEntry) WithFields(fields map[string]interface{}) *LogEntry {
	return &LogEntry{ulog: entry.ulog, Entry: entry.Entry.WithFields(fields)}
}

// WithUUID returns new LogEntry with uuid field in it.
func (entry *LogEntry) WithUUID(uuid string) *LogEntry {
	return &LogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyUUID, uuid)}
}

// WithValidFlag returns new LogEntry with "is valid" field in it.
func (entry *LogEntry) WithValidFlag(isValid bool) *LogEntry {
	return &LogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyIsValid, strconv.FormatBool(isValid))}
}

// WithTime returns new LogEntry with time field in it.
func (entry *LogEntry) WithTime(time time.Time) *LogEntry {
	return &LogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyTime, time.Format(timestampFormat))}
}

// WithTransactionID returns new LogEntry with transaction id field in it.
func (entry *LogEntry) WithTransactionID(tid string) *LogEntry {
	return &LogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyTransactionID, tid)}
}

// WithError returns new LogEntry with error field in it.
func (entry *LogEntry) WithError(err error) *LogEntry {
	return &LogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyError, err)}
}

// WithMonitoringEvent returns new LogEntry with monitoring event fields in it.
// The monitoring event fields are boolean valued monitoring event, event name and content type.
func (entry *LogEntry) WithMonitoringEvent(eventName, tid, contentType string) *LogEntry {
	e := entry.WithFields(
		map[string]interface{}{
			entry.ulog.keyConf.KeyMonitoringEvent: "true",
			entry.ulog.keyConf.KeyEventName:       eventName,
			entry.ulog.keyConf.KeyContentType:     contentType,
		})
	return e.WithTransactionID(tid)
}

// WithCategorisedEvent returns new LogEntry with categorised event fields in it.
// The categorised event fields are event name, event category, event message.
func (entry *LogEntry) WithCategorisedEvent(eventName, eventCategory, eventMsg, tid string) *LogEntry {
	e := entry.WithFields(
		map[string]interface{}{
			entry.ulog.keyConf.KeyEventName:     eventName,
			entry.ulog.keyConf.KeyEventCategory: eventCategory,
			entry.ulog.keyConf.KeyEventMsg:      eventMsg,
		})
	return e.WithTransactionID(tid)
}
