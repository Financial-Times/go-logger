package logger

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func (ulog *UPPLogger) WithError(err error) LogEntry {
	return &logEntry{ulog, ulog.Logger.WithField(ulog.keyConf.KeyError, err)}
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (ulog *UPPLogger) WithField(key string, value interface{}) LogEntry {
	return &logEntry{ulog, ulog.Logger.WithField(key, value)}
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func (ulog *UPPLogger) WithFields(fields map[string]interface{}) LogEntry {
	return &logEntry{ulog, ulog.Logger.WithFields(fields)}
}

// WithTransactionID creates an entry from the standard logger and adds transaction_id field to it.
func (ulog *UPPLogger) WithTransactionID(tid string) LogEntry {
	return ulog.WithField(ulog.keyConf.KeyTransactionID, tid)
}

// WithMonitoringEvent creates an entry from the standard logger and adds monitoring event fields to it.
// The monitoring event fields are "monitoring_event", "event" and "content_type".
func (ulog *UPPLogger) WithMonitoringEvent(eventName, tid, contentType string) LogEntry {
	e := &logEntry{
		ulog,
		ulog.WithField(ulog.keyConf.KeyMonitoringEvent, "true").
			WithField(ulog.keyConf.KeyEventName, eventName).
			WithField(ulog.keyConf.KeyContentType, contentType),
	}
	return e.WithTransactionID(tid)
}

// WithMonitoringEvent creates an entry from the standard logger and adds categorised event fields to it.
// The added fields are "monitoring_event", "event" and "content_type".
func (ulog *UPPLogger) WithCategorisedEvent(eventName, eventCategory, eventMsg, tid string) LogEntry {
	e := &logEntry{
		ulog,
		ulog.WithField(ulog.keyConf.KeyEventName, eventName).
			WithField(ulog.keyConf.KeyEventCategory, eventCategory).
			WithField(ulog.keyConf.KeyEventMsg, eventMsg),
	}
	return e.WithTransactionID(tid)
}
