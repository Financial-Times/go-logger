package logger

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// logEntry is wrapper around logrus.Entry with a few methods adding UPP specific keys to the log entries.
// logEntry is the final or intermediate logging entry. It's finally logged when Debug, Info,
// Warn, Error, Fatal or Panic is called on it.
type logEntry struct {
	ulog *UPPLogger
	*logrus.Entry
}

type LogEntry interface {
	logrus.FieldLogger
	WithUUID(uuid string) LogEntry
	WithValidFlag(isValid bool) LogEntry
	WithTime(time time.Time) LogEntry
	WithTransactionID(tid string) LogEntry
}

func (entry *logEntry) WithUUID(uuid string) LogEntry {
	return &logEntry{ulog: entry.ulog, Entry: entry.WithField(entry.ulog.keyConf.KeyUUID, uuid)}
}

func (entry *logEntry) WithValidFlag(isValid bool) LogEntry {
	return &logEntry{ulog: entry.ulog, Entry: entry.WithField(entry.ulog.keyConf.KeyIsValid, strconv.FormatBool(isValid))}
}

func (entry *logEntry) WithTime(time time.Time) LogEntry {
	return &logEntry{ulog: entry.ulog, Entry: entry.WithField(entry.ulog.keyConf.KeyTime, time.Format(timestampFormat))}
}

func (entry *logEntry) WithTransactionID(tid string) LogEntry {
	return &logEntry{ulog: entry.ulog, Entry: entry.WithField(entry.ulog.keyConf.KeyTransactionID, tid)}
}
