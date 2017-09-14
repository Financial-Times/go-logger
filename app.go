package logger

import (
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

type appLogger struct {
	*log.Logger
	serviceName string
}

type logEntry struct {
	*log.Entry
}

type LogEntry interface {
	log.FieldLogger
	WithUUID(uuid string) LogEntry
	WithValidFlag(isValid bool) LogEntry
	WithTime(time time.Time) LogEntry
}

const (
	serviceStartedEvent = "service_started"
	timestampFormat     = time.RFC3339Nano
)

var logger *appLogger

func InitLogger(serviceName string, logLevel string) {
	parsedLogLevel, err := log.ParseLevel(logLevel)
	if err != nil {
		log.WithFields(log.Fields{"logLevel": logLevel, "err": err}).Fatal("Incorrect log level. Using INFO instead.")
		parsedLogLevel = log.InfoLevel
	}
	log.SetLevel(parsedLogLevel)
	logger = &appLogger{NewLogger(), serviceName}
	logger.Formatter = &log.JSONFormatter{DisableTimestamp: true}
}

func InitDefaultLogger(serviceName string) {
	log.SetLevel(log.InfoLevel)
	logger = &appLogger{NewLogger(), serviceName}
	logger.Formatter = &log.JSONFormatter{DisableTimestamp: true}
}

func NewLogger() *log.Logger {
	return log.New()
}

func NewMonitoringEntry(eventName, tid, contentType string) LogEntry {
	return &logEntry{NewEntry(tid).
		WithField("monitoring_event", "true").
		WithField("event", eventName).
		WithField("content_type", contentType)}

}
func NewEntry(tid string) LogEntry {
	return &logEntry{logger.WithFields(log.Fields{
		"@time":          time.Now().Format(timestampFormat),
		"service_name":   logger.serviceName,
		"transaction_id": tid,
	})}
}

func (entry *logEntry) WithUUID(uuid string) LogEntry {
	return &logEntry{entry.WithField("uuid", uuid)}
}

func (entry *logEntry) WithValidFlag(isValid bool) LogEntry {
	return &logEntry{entry.WithField("isValid", strconv.FormatBool(isValid))}
}

func (entry *logEntry) WithTime(time time.Time) LogEntry {
	return &logEntry{entry.WithField("time", time.Format(timestampFormat))}
}

func ServiceStartedEvent(port int) {
	fields := map[string]interface{}{
		"@time":        time.Now().Format(timestampFormat),
		"service_name": logger.serviceName,
		"event":        serviceStartedEvent,
	}
	logger.WithFields(fields).Infof("Service running on port [%d]", port)
}

func Infof(fields map[string]interface{}, message string, args ...interface{}) {
	entry := logger.WithField("service_name", logger.serviceName).WithField("@time", time.Now().Format(timestampFormat)).WithFields(fields)
	if len(args) > 0 {
		entry.Infof(message, args)
	} else {
		entry.Info(message)
	}
}

func Warnf(fields map[string]interface{}, message string, args ...interface{}) {
	entry := logger.WithField("service_name", logger.serviceName).WithField("@time", time.Now().Format(timestampFormat)).WithFields(fields)
	if len(args) > 0 {
		entry.Warnf(message, args)
	} else {
		entry.Warn(message)
	}
}

func Debugf(fields map[string]interface{}, message string, args ...interface{}) {
	entry := logger.WithField("service_name", logger.serviceName).WithField("@time", time.Now().Format(timestampFormat)).WithFields(fields)
	if len(args) > 0 {
		entry.Debugf(message, args)
	} else {
		entry.Debug(message)
	}
}

func Errorf(fields map[string]interface{}, err error, message string, args ...interface{}) {
	entry := logger.WithField("service_name", logger.serviceName).WithField("@time", time.Now().Format(timestampFormat)).WithFields(fields)
	if err != nil {
		entry = entry.WithError(err)
	}
	if len(args) > 0 {
		entry.Errorf(message, args)
	} else {
		entry.Error(message)
	}
}

func Fatalf(fields map[string]interface{}, err error, message string, args ...interface{}) {
	entry := logger.WithField("service_name", logger.serviceName).WithField("@time", time.Now().Format(timestampFormat)).WithFields(fields)
	if err != nil {
		entry = entry.WithError(err)
	}
	if len(args) > 0 {
		entry.Fatalf(message, args)
	} else {
		entry.Fatal(message)
	}
}
