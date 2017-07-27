package logger

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

type appLogger struct {
	log         *logrus.Logger
	serviceName string
}

const (
	serviceStartedEvent = "service_started"
	mappingEvent        = "mapping"
)

var logger *appLogger

func InitLogger(serviceName string, logLevel string) {
	parsedLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{"logLevel": logLevel, "err": err}).Fatal("Incorrect log level. Using INFO instead.")
		parsedLogLevel = logrus.InfoLevel
	}
	logrus.SetLevel(parsedLogLevel)
	log := NewLogger()
	log.Formatter = new(logrus.JSONFormatter)
	logger = &appLogger{log, serviceName}
}

func InitDefaultLogger(serviceName string) {
	logrus.SetLevel(logrus.InfoLevel)
	log := NewLogger()
	log.Formatter = new(logrus.JSONFormatter)
	logger = &appLogger{log, serviceName}
}

func NewLogger() *logrus.Logger {
	return logrus.New()
}

//****************** MONITORING LOGS ******************
func MonitoringEvent(eventName, tid, contentType, message string) {
	logger.log.WithFields(logrus.Fields{
		"event":            eventName,
		"monitoring_event": true,
		"service_name":     logger.serviceName,
		"transaction_id":   tid,
		"content_type":     contentType,
	}).Info(message)
}

func MonitoringEventWithUUID(eventName, tid, uuid, contentType, message string) {
	logger.log.WithFields(logrus.Fields{
		"event":            eventName,
		"monitoring_event": true,
		"transaction_id":   tid,
		"uuid":             uuid,
		"content_type":     contentType,
		"service_name":     logger.serviceName,
	}).Info(message)
}

func MonitoringValidationEvent(tid, uuid, contentType, message string, isValid bool) {
	if isValid {
		logger.log.WithFields(logrus.Fields{
			"event":            mappingEvent,
			"monitoring_event": "true",
			"transaction_id":   tid,
			"uuid":             uuid,
			"content_type":     contentType,
			"service_name":     logger.serviceName,
			"isValid":          strconv.FormatBool(isValid),
		}).Info(message)
	} else {
		logger.log.WithFields(logrus.Fields{
			"event":            mappingEvent,
			"monitoring_event": "true",
			"transaction_id":   tid,
			"uuid":             uuid,
			"content_type":     contentType,
			"service_name":     logger.serviceName,
			"isValid":          strconv.FormatBool(isValid),
		}).Error(message)
	}
}

//****************** SERVICE LOGS ******************
func ServiceStartedEvent(port int) {
	fields := map[string]interface{}{
		"service_name": logger.serviceName,
		"event":        serviceStartedEvent,
	}
	logger.log.WithFields(fields).Infof("Service running on port [%d]", port)
}

func InfoEvent(transactionID string, message string) {
	fields := map[string]interface{}{
		"service_name":   logger.serviceName,
		"transaction_id": transactionID,
	}
	logger.log.WithFields(fields).Info(message)
}

func InfoEventWithUUID(transactionID string, contentUUID string, message string) {
	fields := map[string]interface{}{
		"service_name":   logger.serviceName,
		"transaction_id": transactionID,
	}
	if contentUUID != "" {
		fields["uuid"] = contentUUID
	}
	logger.log.WithFields(fields).Info(message)
}

func WarnEvent(transactionID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": logger.serviceName,
	}
	if err != nil {
		fields["error"] = err
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	logger.log.WithFields(fields).Warn(message)
}

func WarnEventWithUUID(transactionID string, contentUUID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": logger.serviceName,
	}
	if err != nil {
		fields["error"] = err
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	if contentUUID != "" {
		fields["uuid"] = contentUUID
	}
	logger.log.WithFields(fields).Warn(message)
}

func ErrorEvent(transactionID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": logger.serviceName,
		"error":        err,
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	logger.log.WithFields(fields).Error(message)
}

func ErrorEventWithUUID(transactionID string, contentUUID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": logger.serviceName,
		"error":        err,
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	if contentUUID != "" {
		fields["uuid"] = contentUUID
	}
	logger.log.WithFields(fields).Error(message)
}

func FatalEvent(message string, err error) {
	fields := map[string]interface{}{
		"service_name": logger.serviceName,
		"error":        err,
	}
	logger.log.WithFields(fields).Fatal(message)
}

//****************** SERVICE general structured LOGS ******************
func Infof(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = logger.serviceName
	logger.log.WithFields(fields).Infof(message, args)
}

func Warnf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = logger.serviceName
	logger.log.WithFields(fields).Warnf(message, args)
}

func Debugf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = logger.serviceName
	logger.log.WithFields(fields).Debugf(message, args)
}

func Errorf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = logger.serviceName
	logger.log.WithFields(fields).Errorf(message, args)
}

func Fatalf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = logger.serviceName
	logger.log.WithFields(fields).Fatalf(message, args)
}
