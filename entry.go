package logger

//go:generate moq -out entry_mock.go . LogEntry

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

type LogPrinter interface {
	Debug(args ...interface{})
	Print(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Warning(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Printf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Warningf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Debugln(args ...interface{})
	Infoln(args ...interface{})
	Println(args ...interface{})
	Warnln(args ...interface{})
	Warningln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}

type LogEntry interface {
	LogPrinter
	Logger
}

// LogEntry is wrapper around logrus.Entry with a few methods adding UPP specific keys to the log entries.
// LogEntry is the final or intermediate logging entry. It's finally logged when Debug, Info,
// Warn, Error, Fatal or Panic is called on it.
type UPPLogEntry struct {
	ulog *UPPLogger
	*logrus.Entry
}

// WithField for LogEntry is wrapper around WithField of logrus.Entry
// that let us return our object when chaining the calls to WithField.
func (entry *UPPLogEntry) WithField(key string, value interface{}) LogEntry {
	return &UPPLogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(key, value)}
}

// WithFields for LogEntry is wrapper around WithFields of logrus.Entry
// that let us return our object when chaining the calls to WithFields.
func (entry *UPPLogEntry) WithFields(fields map[string]interface{}) LogEntry {
	return &UPPLogEntry{ulog: entry.ulog, Entry: entry.Entry.WithFields(fields)}
}

// WithUUID returns new LogEntry with uuid field in it.
func (entry *UPPLogEntry) WithUUID(uuid string) LogEntry {
	return &UPPLogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyUUID, uuid)}
}

// WithValidFlag returns new LogEntry with "is valid" field in it.
func (entry *UPPLogEntry) WithValidFlag(isValid bool) LogEntry {
	return &UPPLogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyIsValid, strconv.FormatBool(isValid))}
}

// WithTime returns new LogEntry with time field in it.
func (entry *UPPLogEntry) WithTime(time time.Time) LogEntry {
	return &UPPLogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyTime, time.Format(timestampFormat))}
}

// WithTransactionID returns new LogEntry with transaction id field in it.
func (entry *UPPLogEntry) WithTransactionID(tid string) LogEntry {
	return &UPPLogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyTransactionID, tid)}
}

// WithError returns new LogEntry with error field in it.
func (entry *UPPLogEntry) WithError(err error) LogEntry {
	return &UPPLogEntry{ulog: entry.ulog, Entry: entry.Entry.WithField(entry.ulog.keyConf.KeyError, err)}
}

// WithMonitoringEvent returns new LogEntry with monitoring event fields in it.
// The monitoring event fields are boolean valued monitoring event, event name and content type.
func (entry *UPPLogEntry) WithMonitoringEvent(eventName, tid, contentType string) LogEntry {
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
func (entry *UPPLogEntry) WithCategorisedEvent(eventName, eventCategory, eventMsg, tid string) LogEntry {
	e := entry.WithFields(
		map[string]interface{}{
			entry.ulog.keyConf.KeyEventName:     eventName,
			entry.ulog.keyConf.KeyEventCategory: eventCategory,
			entry.ulog.keyConf.KeyEventMsg:      eventMsg,
		})
	return e.WithTransactionID(tid)
}
