package go_logger

import (
	"github.com/Sirupsen/logrus"
)

type AppLogger struct {
	log         *logrus.Logger
	serviceName string
}

const (
	serviceStartedEvent = "service_started"
	mappingEvent        = "mapping"
)

func NewLogger(serviceName string) *AppLogger {
	logrus.SetLevel(logrus.InfoLevel)
	log := logrus.New()
	log.Formatter = new(logrus.JSONFormatter)
	return &AppLogger{log, serviceName}
}

//****************** MONITORING LOGS ******************
func (appLogger *AppLogger) MonitoringEvent(eventName, tid, contentType, message string) {
	appLogger.log.WithFields(logrus.Fields{
		"event":            eventName,
		"monitoring_event": true,
		"service_name":     appLogger.serviceName,
		"transaction_id":   tid,
		"content_type":     contentType,
	}).Info(message)
}

func (appLogger *AppLogger) MonitoringEventWithUUID(eventName, tid, uuid, contentType, message string) {
	appLogger.log.WithFields(logrus.Fields{
		"event":            eventName,
		"monitoring_event": true,
		"transaction_id":   tid,
		"uuid":             uuid,
		"content_type":     contentType,
		"service_name":     appLogger.serviceName,
	}).Info(message)
}

func (appLogger *AppLogger) MonitoringValidationEvent(tid, uuid, contentType, message string, isValid bool) {
	if isValid {
		appLogger.log.WithFields(logrus.Fields{
			"event":            mappingEvent,
			"monitoring_event": true,
			"transaction_id":   tid,
			"uuid":             uuid,
			"content_type":     contentType,
			"service_name":     appLogger.serviceName,
			"isValid":          isValid,
		}).Info(message)
	} else {
		appLogger.log.WithFields(logrus.Fields{
			"event":            mappingEvent,
			"monitoring_event": true,
			"transaction_id":   tid,
			"uuid":             uuid,
			"content_type":     contentType,
			"service_name":     appLogger.serviceName,
			"isValid":          isValid,
		}).Error(message)
	}
}

//****************** SERVICE LOGS ******************
func (appLogger *AppLogger) ServiceStartedEvent(port int) {
	fields := map[string]interface{}{
		"service_name": appLogger.serviceName,
		"event":        serviceStartedEvent,
	}
	appLogger.log.WithFields(fields).Infof("Service running on port [%d]", port)
}

func (appLogger *AppLogger) InfoEvent(transactionID string, message string) {
	fields := map[string]interface{}{
		"service_name":   appLogger.serviceName,
		"transaction_id": transactionID,
	}
	appLogger.log.WithFields(fields).Info(message)
}

func (appLogger *AppLogger) InfoEventWithUUID(transactionID string, contentUUID string, message string) {
	fields := map[string]interface{}{
		"service_name":   appLogger.serviceName,
		"transaction_id": transactionID,
	}
	if contentUUID != "" {
		fields["uuid"] = contentUUID
	}
	appLogger.log.WithFields(fields).Info(message)
}

func (appLogger *AppLogger) WarnEvent(transactionID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.serviceName,
	}
	if err != nil {
		fields["error"] = err
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	appLogger.log.WithFields(fields).Warn(message)
}

func (appLogger *AppLogger) WarnEventWithUUID(transactionID string, contentUUID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.serviceName,
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
	appLogger.log.WithFields(fields).Warn(message)
}

func (appLogger *AppLogger) ErrorEvent(transactionID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.serviceName,
		"error":        err,
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	appLogger.log.WithFields(fields).Error(message)
}

func (appLogger *AppLogger) ErrorEventWithUUID(transactionID string, contentUUID string, message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.serviceName,
		"error":        err,
	}
	if transactionID != "" {
		fields["transaction_id"] = transactionID
	}
	if contentUUID != "" {
		fields["uuid"] = contentUUID
	}
	appLogger.log.WithFields(fields).Error(message)
}

func (appLogger *AppLogger) FatalEvent(message string, err error) {
	fields := map[string]interface{}{
		"service_name": appLogger.serviceName,
		"error":        err,
	}
	appLogger.log.WithFields(fields).Fatal(message)
}

//****************** SERVICE general structured LOGS ******************
func (appLogger *AppLogger) Infof(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = appLogger.serviceName
	appLogger.log.WithFields(fields).Infof(message, args)
}

func (appLogger *AppLogger) Warnf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = appLogger.serviceName
	appLogger.log.WithFields(fields).Warnf(message, args)
}

func (appLogger *AppLogger) Errorf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = appLogger.serviceName
	appLogger.log.WithFields(fields).Errorf(message, args)
}

func (appLogger *AppLogger) Fatalf(fields map[string]interface{}, message string, args ...interface{}) {
	fields["service_name"] = appLogger.serviceName
	appLogger.log.WithFields(fields).Fatalf(message, args)
}
