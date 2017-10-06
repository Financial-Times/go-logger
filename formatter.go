package logger

import (
	"encoding/json"
	"errors"
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

func newFTJSONFormatter() *ftJSONFormatter {
	return &ftJSONFormatter{}
}

func (f *ftJSONFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	if f.serviceName == "" {
		return []byte{}, errors.New("logger is not initialised - please use InitLogger or InitDefaultLogger function")
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

	if _, found := data[fieldKeyTime]; !found {
		data[fieldKeyTime] = entry.Time.Format(timestampFormat)
	}

	data[logrus.FieldKeyMsg] = entry.Message
	data[logrus.FieldKeyLevel] = entry.Level.String()
	data[fieldKeyServiceName] = f.serviceName

	serialized, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Failed to marshal fields to JSON, %v", err)
	}
	return append(serialized, '\n'), nil
}
