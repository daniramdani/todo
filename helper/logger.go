package helper

import (
	log "github.com/sirupsen/logrus"
)

const (
	TOPIC = "user-service-log"
	LogTag = "user-service"
)

func LogContext(c string, s string) *log.Entry {
	return log.WithFields(log.Fields{
		"topic":   TOPIC,
		"context": c,
		"scope":   s,
	})
}

func Capture(level log.Level, err error, context string, scope string) {
	log.SetFormatter(&log.JSONFormatter{})

	entry := LogContext(context, scope)
	switch level {
	case log.ErrorLevel:
		entry.Error(err.Error())
	case log.PanicLevel:
		entry.Panic(err.Error())
	}
}
