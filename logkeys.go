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
// if there are key names missing from the input conf, the default key names are used
func GetFullKeyNameConfig(conf KeyNamesConfig) *KeyNamesConfig {
	defaultConfig := GetDefaultKeyNamesConfig()
	resConfig := conf

	if keyLogLevel := conf.KeyLogLevel; keyLogLevel == "" {
		resConfig.KeyLogLevel = defaultConfig.KeyLogLevel
	}
	if keyMsg := conf.KeyMsg; keyMsg == "" {
		resConfig.KeyMsg = defaultConfig.KeyMsg
	}
	if keyError := conf.KeyError; keyError == "" {
		resConfig.KeyError = defaultConfig.KeyError
	}
	if keyTime := conf.KeyTime; keyTime == "" {
		resConfig.KeyTime = defaultConfig.KeyTime
	}
	if keyServiceName := conf.KeyServiceName; keyServiceName == "" {
		resConfig.KeyServiceName = defaultConfig.KeyServiceName
	}
	if keyTransactionID := conf.KeyTransactionID; keyTransactionID == "" {
		resConfig.KeyTransactionID = defaultConfig.KeyTransactionID
	}
	if keyUUID := conf.KeyUUID; keyUUID == "" {
		resConfig.KeyUUID = defaultConfig.KeyUUID
	}
	if keyIsValid := conf.KeyIsValid; keyIsValid == "" {
		resConfig.KeyIsValid = defaultConfig.KeyIsValid
	}
	if keyEventName := conf.KeyEventName; keyEventName == "" {
		resConfig.KeyEventName = defaultConfig.KeyEventName
	}
	if keyMonitoringEvent := conf.KeyMonitoringEvent; keyMonitoringEvent == "" {
		resConfig.KeyMonitoringEvent = defaultConfig.KeyMonitoringEvent
	}
	if keyContentType := conf.KeyContentType; keyContentType == "" {
		resConfig.KeyContentType = defaultConfig.KeyContentType
	}
	if keyEventCategory := conf.KeyEventCategory; keyEventCategory == "" {
		resConfig.KeyEventCategory = defaultConfig.KeyEventCategory
	}
	if keyEventMsg := conf.KeyEventMsg; keyEventMsg == "" {
		resConfig.KeyEventMsg = defaultConfig.KeyEventMsg
	}
	return &resConfig
}
