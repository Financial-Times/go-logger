package logger

import "github.com/sirupsen/logrus"

const (
	DefaultKeyLogLevel = logrus.FieldKeyLevel
	DefaultKeyMsg      = logrus.FieldKeyMsg
	DefaultKeyError    = "error"
	DefaultKeyTime     = "@time"

	DefaultKeyServiceName   = "service_name"
	DefaultKeyTransactionID = "transaction_id"
	DefaultKeyUUID          = "uuid"
	DefaultKeyIsValid       = "isValid"

	DefaultKeyEventName       = "event"
	DefaultKeyMonitoringEvent = "monitoring_event"
	DefaultKeyContentType     = "content_type"
	DefaultKeyEventCategory   = "event_category"
	DefaultKeyEventMsg        = "event_msg"
)

type KeyNamesConfig struct {
	KeyLogLevel string
	KeyMsg      string
	KeyError    string
	KeyTime     string

	KeyServiceName   string
	KeyTransactionID string
	KeyUUID          string
	KeyIsValid       string

	KeyEventName       string
	KeyMonitoringEvent string
	KeyContentType     string
	KeyEventCategory   string
	KeyEventMsg        string
}

func GetDefaultKeyNamesConfig() *KeyNamesConfig {
	return &KeyNamesConfig{
		KeyLogLevel:        DefaultKeyLogLevel,
		KeyMsg:             DefaultKeyMsg,
		KeyError:           DefaultKeyError,
		KeyTime:            DefaultKeyTime,
		KeyServiceName:     DefaultKeyServiceName,
		KeyTransactionID:   DefaultKeyTransactionID,
		KeyUUID:            DefaultKeyUUID,
		KeyIsValid:         DefaultKeyIsValid,
		KeyEventName:       DefaultKeyEventName,
		KeyMonitoringEvent: DefaultKeyMonitoringEvent,
		KeyContentType:     DefaultKeyContentType,
		KeyEventCategory:   DefaultKeyEventCategory,
		KeyEventMsg:        DefaultKeyEventMsg,
	}
}

// GetFullKeyNameConfig returns KeyNamesConfig that has all key names from the input conf and
// if there are key names missing from the input conf, the default key names are used.
func GetFullKeyNameConfig(conf KeyNamesConfig) *KeyNamesConfig {
	defaultConfig := GetDefaultKeyNamesConfig()

	if conf.KeyLogLevel == "" {
		conf.KeyLogLevel = defaultConfig.KeyLogLevel
	}
	if conf.KeyMsg == "" {
		conf.KeyMsg = defaultConfig.KeyMsg
	}
	if conf.KeyError == "" {
		conf.KeyError = defaultConfig.KeyError
	}
	if conf.KeyTime == "" {
		conf.KeyTime = defaultConfig.KeyTime
	}
	if conf.KeyServiceName == "" {
		conf.KeyServiceName = defaultConfig.KeyServiceName
	}
	if conf.KeyTransactionID == "" {
		conf.KeyTransactionID = defaultConfig.KeyTransactionID
	}
	if conf.KeyUUID == "" {
		conf.KeyUUID = defaultConfig.KeyUUID
	}
	if conf.KeyIsValid == "" {
		conf.KeyIsValid = defaultConfig.KeyIsValid
	}
	if conf.KeyEventName == "" {
		conf.KeyEventName = defaultConfig.KeyEventName
	}
	if conf.KeyMonitoringEvent == "" {
		conf.KeyMonitoringEvent = defaultConfig.KeyMonitoringEvent
	}
	if conf.KeyContentType == "" {
		conf.KeyContentType = defaultConfig.KeyContentType
	}
	if conf.KeyEventCategory == "" {
		conf.KeyEventCategory = defaultConfig.KeyEventCategory
	}
	if conf.KeyEventMsg == "" {
		conf.KeyEventMsg = defaultConfig.KeyEventMsg
	}
	return &conf
}
