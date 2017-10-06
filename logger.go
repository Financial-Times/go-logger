package logger

import (
	"github.com/sirupsen/logrus"
	"time"
)

const (
	serviceStartedEvent = "service_started"
	timestampFormat     = time.RFC3339Nano
)

var log = logrus.New()
var formatter = newFTJSONFormatter()

func init() {
	log.Formatter = formatter
}

func InitLogger(serviceName string, logLevel string) {
	parsedLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		log.WithFields(logrus.Fields{"logLevel": logLevel, "err": err}).Fatal("Incorrect log level. Using INFO instead.")
		parsedLogLevel = logrus.InfoLevel
	}
	log.SetLevel(parsedLogLevel)
	formatter.serviceName = serviceName
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
