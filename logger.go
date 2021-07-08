package logger

import (
	"time"

	"github.com/sirupsen/logrus"
)

const (
	serviceStartedEvent = "service_started"
)

type Logger interface {
	LogPrinter
	WithField(string, interface{}) LogEntry
	WithFields(map[string]interface{}) LogEntry
	WithUUID(string) LogEntry
	WithValidFlag(bool) LogEntry
	WithTime(time.Time) LogEntry
	WithTransactionID(string) LogEntry
	WithError(error) LogEntry
	WithMonitoringEvent(string, string, string) LogEntry
	WithCategorisedEvent(string, string, string, string) LogEntry
	LogServiceStartedEvent(int)
}

// UPPLogger wraps logrus logger providing the same functionality as logrus with a few UPP specifics.
type UPPLogger struct {
	*logrus.Logger
	keyConf *KeyNamesConfig
}

// NewUPPLogger initializes UPP logger with structured logging format.
func NewUPPLogger(serviceName string, logLevel string, kconf ...KeyNamesConfig) Logger {
	keyConf := GetDefaultKeyNamesConfig()
	if len(kconf) > 0 {
		keyConf = GetFullKeyNameConfig(kconf[0])
	}

	logrusLog := logrus.New()
	formatter := newFTJSONFormatter(serviceName, keyConf)
	logrusLog.Formatter = formatter

	parsedLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrusLog.WithField("logLevel", logLevel).WithError(err).Error("Incorrect log level. Using INFO instead.")
		parsedLogLevel = logrus.InfoLevel
	}
	logrusLog.SetLevel(parsedLogLevel)

	return &UPPLogger{Logger: logrusLog, keyConf: keyConf}
}

// NewUPPInfoLogger initializes UPPLogger with log level INFO.
func NewUPPInfoLogger(serviceName string, kconf ...KeyNamesConfig) Logger {
	return NewUPPLogger(serviceName, logrus.InfoLevel.String(), kconf...)
}

// NewUnstructuredLogger initializes plain logrus log without taking into account UPP log formatting.
func NewUnstructuredLogger() Logger {
	return &UPPLogger{Logger: logrus.New(), keyConf: GetDefaultKeyNamesConfig()}
}

// @todo - this is not used; maybe remove it
// LogServiceStartedEvent logs service started event with level INFO.
func (ulog *UPPLogger) LogServiceStartedEvent(port int) {
	fields := map[string]interface{}{
		ulog.keyConf.KeyEventName: serviceStartedEvent,
	}
	ulog.WithFields(fields).Infof("Service running on port [%d]", port)
}
