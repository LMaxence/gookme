package logging

import (
	"os"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func readLevelFromEnv() log.Level {
	level, err := log.ParseLevel(os.Getenv("LOG_LEVEL"))

	if err != nil {
		return log.InfoLevel
	}

	return level
}

func readDebugModeFromEnv() bool {
	_, isSet := os.LookupEnv("DEBUG")
	return isSet
}

type Logger struct {
	fields log.Fields
}

func NewLogger(name string) *Logger {
	level := readLevelFromEnv()
	log.SetLevel(level)

	log.SetFormatter(&log.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	return &Logger{
		fields: log.Fields{
			"name": name,
		},
	}
}

func (l *Logger) applyFields() *logrus.Entry {
	if !readDebugModeFromEnv() {
		return log.WithFields(log.Fields{})
	}
	return log.WithFields(l.fields)
}

func (l *Logger) WithFields(fields map[string]string) *Logger {
	newFields := log.Fields{}

	for key, value := range l.fields {
		newFields[key] = value
	}

	for key, value := range fields {
		newFields[key] = value
	}

	return &Logger{
		fields: newFields,
	}
}

func (l *Logger) Debug(args ...interface{}) {
	l.applyFields().Debug(args...)
}

func (l *Logger) Debugf(format string, args ...interface{}) {
	l.applyFields().Debugf(format, args...)
}

func (l *Logger) Trace(args ...interface{}) {
	l.applyFields().Trace(args...)
}

func (l *Logger) Tracef(format string, args ...interface{}) {
	l.applyFields().Tracef(format, args...)
}

func (l *Logger) Info(args ...interface{}) {
	l.applyFields().Info(args...)
}

func (l *Logger) Infof(format string, args ...interface{}) {
	l.applyFields().Infof(format, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.applyFields().Warn(args...)
}

func (l *Logger) Warnf(format string, args ...interface{}) {
	l.applyFields().Warnf(format, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.applyFields().Error(args...)
}

func (l *Logger) Errorf(format string, args ...interface{}) {
	l.applyFields().Errorf(format, args...)
}
