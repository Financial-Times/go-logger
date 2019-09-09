package logger

import (
	"strconv"
	"time"

	"github.com/sirupsen/logrus"
)

// logEntry is wrapper around logrus.Entry adding a few methods adding UPP specific keys to the log entries.
type logEntry struct {
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
	return &logEntry{entry.WithField("uuid", uuid)}
}

func (entry *logEntry) WithValidFlag(isValid bool) LogEntry {
	return &logEntry{entry.WithField("isValid", strconv.FormatBool(isValid))}
}

func (entry *logEntry) WithTime(time time.Time) LogEntry {
	return &logEntry{entry.WithField(fieldKeyTime, time.Format(timestampFormat))}
}

func (entry *logEntry) WithTransactionID(tid string) LogEntry {
	return &logEntry{entry.WithField("transaction_id", tid)}
}
