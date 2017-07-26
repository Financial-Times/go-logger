package go_logger

import (
	"github.com/Sirupsen/logrus"
)

type AppLogger struct {
	Log         *logrus.Logger
	ServiceName string
}

const (
	serviceStartedEvent = "service_started"
	mappingEvent        = "mapping"
)

func NewConfLogger(serviceName string, logLevel string) *AppLogger {
	parsedLogLevel, err := logrus.ParseLevel(logLevel)
	if err != nil {
		logrus.WithFields(logrus.Fields{"logLevel": logLevel, "err": err}).Fatal("Incorrect log level. Using INFO instead.")
		parsedLogLevel = logrus.InfoLevel
	}
	logrus.SetLevel(parsedLogLevel)
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	return &AppLogger{log, serviceName}
}

func NewLogger(serviceName string) *AppLogger {
	logrus.SetLevel(logrus.InfoLevel)
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	return &AppLogger{log, serviceName}
}

//****************** MONITORING LOGS ******************
func (appLogger *AppLogger) MonitoringEvent(eventName, tid, contentType, message string) {
	appLogger.Log.WithFields(logrus.Fields{
		"event":            eventName,
		"monitoring_event": true,
		"service_name":     appLogger.ServiceName,
		"transaction_id":   tid,
		"content_type":     contentType,
	}).Info(message)
}

func (appLogger *AppLogger) MonitoringEventWithUUID(eventName, tid, uuid, contentType, message string) {
	appLogger.Log.WithFields(logrus.Fields{
		"event":            eventName,
		"monitoring_event": true,
		"transaction_id":   tid,
		"uuid":             uuid,
		"content_type":     contentType,
		"service_name":     appLogger.ServiceName,
	}).Info(message)
}

func (appLogger *AppLogger) MonitoringValidationEvent(tid, uuid, contentType, message string, isValid bool) {
	if isValid {
		appLogger.Log.WithFields(logrus.Fields{
			"event":            mappingEvent,
			"monitoring_event": true,
			"transaction_id":   tid,
			"uuid":             uuid,
			"content_type":     contentType,
			"service_name":     appLogger.ServiceName,
			"isValid":          isValid,
		}).Info(message)
	} else {
		appLogger.Log.WithFields(logrus.Fields{
			"event":            mappingEvent,
			"monitoring_event": true,
			"transaction_id":   tid,
			"uuid":             uuid,
			"content_type":     contentType,
			"service_name":     appLogger.ServiceName,
			"isValid":          isValid,
		}).Error(message)
	}
}

//****************** SERVICE LOGS ******************
func (appLogger *AppLogger) ServiceStartedEvent(port int) {
	fields := map[string]interface{}{
		"service_name": appLogger.ServiceName,
		"event":        serviceStartedEvent,
	}
	appLogger.Log.WithFields(fields).Infof("Service running on port [%d]", port)
}

func (appLogger *AppLogger) InfoEvent(transactionID string, message string) {
	fields := map[string]interface{}{
		"service_name":   appLogger.ServiceName,
		"transaction_id": transactionID,
	}
	appLogger.Log.WithFields(fields).Info(message)
}

func (appLogger *AppLogger) InfoEventWithUUID(transactionID string, contentUUID string, message string) {
	fields := map[string]interface{}{
		"service_name":   appLogger.ServiceName,
		"transaction_id": transactionID,
	}
	if contentUUID != "" {
		fields["uuid"] = contentUUID
	}
	appLogger.Log.WithFields(fields).Info(message)
}

func (appLogger *AppLogger) WarnEvent(transactionID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.ServiceName,
	}
	if err != nil {
		fields["error"] = err
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	appLogger.Log.WithFields(fields).Warn(message)
}

func (appLogger *AppLogger) WarnEventWithUUID(transactionID string, contentUUID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.ServiceName,
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
	appLogger.Log.WithFields(fields).Warn(message)
}

func (appLogger *AppLogger) ErrorEvent(transactionID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.ServiceName,
		"error":        err,
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	appLogger.Log.WithFields(fields).Error(message)
}

func (appLogger *AppLogger) ErrorEventWithUUID(transactionID string, contentUUID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.ServiceName,
		"error":        err,
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	if contentUUID != "" {
		fields["uuid"] = contentUUID
	}
	appLogger.Log.WithFields(fields).Error(message)
}

func (appLogger *AppLogger) FatalEvent(message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.ServiceName,
		"error":        err,
	}
	appLogger.Log.WithFields(fields).Fatal(message)
}

//****************** SERVICE general structured LOGS ******************
func (appLogger *AppLogger) Infof(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = appLogger.ServiceName
	appLogger.Log.WithFields(fields).Infof(message, args)
}

func (appLogger *AppLogger) Warnf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = appLogger.ServiceName
	appLogger.Log.WithFields(fields).Warnf(message, args)
}

func (appLogger *AppLogger) Errorf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = appLogger.ServiceName
	appLogger.Log.WithFields(fields).Errorf(message, args)
}

func (appLogger *AppLogger) Fatalf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = appLogger.ServiceName
	appLogger.Log.WithFields(fields).Fatalf(message, args)
}
