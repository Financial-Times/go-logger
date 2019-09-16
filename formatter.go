package logger

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

const timestampFormat = time.RFC3339Nano

// ftJSONFormatter formats the logs in JSON format.
// It always includes "msg", "level" and "service_name" fields for each log entry.
// If there is time field in the log entry, ftJSONFormatter logs it in time.RFC3339Nano format.
type ftJSONFormatter struct {
	serviceName string
	keyConf     *KeyNamesConfig
}

func newFTJSONFormatter(serviceName string, keyConf *KeyNamesConfig) *ftJSONFormatter {
	return &ftJSONFormatter{serviceName: serviceName, keyConf: keyConf}
}

func (f *ftJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.serviceName == "" {
		return []byte{}, errors.New("UPP log formatter is not initialised with service name")
	}

	data := make(logrus.Fields)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	if _, found := data[f.keyConf.KeyTime]; !found {
		data[f.keyConf.KeyTime] = entry.Time.Format(timestampFormat)
	}

	data[f.keyConf.KeyMsg] = entry.Message
	data[f.keyConf.KeyLogLevel] = entry.Level.String()
	data[f.keyConf.KeyServiceName] = f.serviceName

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
