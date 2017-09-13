package logger

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
)

const (
	fieldKeyTime        = "@time"
	fieldKeyServiceName = "service_name"
)

type ftJSONFormatter struct {
	serviceName string
}

func newFTJSONFormatter(serviceName string) logrus.Formatter {
	return &ftJSONFormatter{
		serviceName: serviceName,
	}
}

func (f *ftJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	data := make(logrus.Fields, len(entry.Data)+4)
	for k, v := range entry.Data {
		switch v := v.(type) {
		case error:
			// Otherwise errors are ignored by `encoding/json`
			data[k] = v.Error()
		default:
			data[k] = v
		}
	}

	data[fieldKeyTime] = entry.Time.Format(timestampFormat)
	data[logrus.FieldKeyMsg] = entry.Message
	data[logrus.FieldKeyLevel] = entry.Level.String()
	data[fieldKeyServiceName] = f.serviceName

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
