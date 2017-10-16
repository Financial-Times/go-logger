package logger

import (
	"github.com/sirupsen/logrus"
)

const (
	serviceStartedEvent = "service_started"
)

var log = logrus.New()
var formatter = newFTJSONFormatter()

func init() {
	log.Formatter = formatter
}

func InitLogger(serviceName string, logLevel string) {
	formatter.serviceName = serviceName
	parsedLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.WithField("logLevel", logLevel).WithError(err).Error("Incorrect log level. Using INFO instead.")
		parsedLogLevel = logrus.InfoLevel
	}
	log.SetLevel(parsedLogLevel)
}

func InitDefaultLogger(serviceName string) {
	InitLogger(serviceName, logrus.InfoLevel.String())
}

func ServiceStartedEvent(port int) {
	fields := map[string]interface{}{
		"event": serviceStartedEvent,
	}
	log.WithFields(fields).Infof("Service running on port [%d]", port)
}

func Logger() *logrus.Logger {
	return log
}
