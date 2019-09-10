package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	serviceStartedEvent = "service_started"
)

// UPPLogger wraps logrus logger providing the same functionality as logrus with a few UPP specifics.
type UPPLogger struct {
	*logrus.Logger
	keyConf *KeyNamesConfig
}

// NewUPPLogger initializes UPP logger with structured logging format.
func NewUPPLogger(serviceName string, logLevel string, kconf ...KeyNamesConfig) *UPPLogger {
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
func NewUPPInfoLogger(serviceName string, kconf ...KeyNamesConfig) *UPPLogger {
	return NewUPPLogger(serviceName, logrus.InfoLevel.String(), kconf...)
}

// NewUnstructuredLogger initializes plain logrus log without taking into account UPP log formatting.
func NewUnstructuredLogger() *UPPLogger {
	return &UPPLogger{Logger: logrus.New(), keyConf: GetDefaultKeyNamesConfig()}
}

// LogServiceStartedEvent logs service started event with level INFO.
func (ulog *UPPLogger) LogServiceStartedEvent(port int) {
	fields := map[string]interface{}{
		ulog.keyConf.KeyEventName: serviceStartedEvent,
	}
	ulog.WithFields(fields).Infof("Service running on port [%d]", port)
}
