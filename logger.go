package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	serviceStartedEvent = "service_started"
)

// UPPLogger wraps logrus logger providing the same functionality as logrus with a few UPP specifics
type UPPLogger struct{
	*logrus.Logger
}

// NewUnstructuredLogger initializes plain logrus log without taking into account UPP log formatting
func NewUnstructuredLogger() *UPPLogger {
	return &UPPLogger{logrus.New()}
}

// NewUPPLogger initializes UPP logger with structured logging format
func NewUPPLogger(serviceName string, logLevel string) *UPPLogger{
	logrusLog := logrus.New()
	formatter := newFTJSONFormatter()
	formatter.serviceName = serviceName

	parsedLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrusLog.WithField("logLevel", logLevel).WithError(err).Error("Incorrect log level. Using INFO instead.")
		parsedLogLevel = logrus.InfoLevel
	}
	logrusLog.SetLevel(parsedLogLevel)
	return &UPPLogger{logrusLog}
}

// NewUPPInfoLogger initializes UPPLogger with log level INFO
func NewUPPInfoLogger(serviceName string) *UPPLogger{
	return NewUPPLogger(serviceName, logrus.InfoLevel.String())
}

// LogServiceStartedEvent logs service started event with level INFO
func (ulog *UPPLogger) LogServiceStartedEvent(port int) {
	fields := map[string]interface{}{
		"event": serviceStartedEvent,
	}
	ulog.WithFields(fields).Infof("Service running on port [%d]", port)
}
