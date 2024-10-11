package logging

import (
	"os"

	logrus "github.com/sirupsen/logrus"
)

// readLevelFromEnv reads the log level from the LOG_LEVEL environment variable
// and returns the corresponding logrus level. If the LOG_LEVEL environment variable
// is not set, the default level is INFO.
func readLevelFromEnv() logrus.Level {
	level, err := logrus.ParseLevel(os.Getenv("LOG_LEVEL"))

	if err != nil {
		return logrus.InfoLevel
	}

	return level
}

func readDebugModeFromEnv() bool {
	_, isSet := os.LookupEnv("DEBUG")
	return isSet
}

type Logger struct {
	fields logrus.Fields
}

// NewLogger creates a new logger with the provided name
func NewLogger(name string) *Logger {
	level := readLevelFromEnv()
	logrus.SetLevel(level)

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp:       true,
		DisableLevelTruncation: true,
		PadLevelText:           true,
	})

	return &Logger{
		fields: logrus.Fields{
			"name": name,
		},
	}
}

// applyFields applies the fields stored in the logger to the logrus entry
func (l *Logger) applyFields() *logrus.Entry {
	if !readDebugModeFromEnv() {
		return logrus.WithFields(logrus.Fields{})
	}
	return logrus.WithFields(l.fields)
}

// WithFields returns a new logger with the provided fields
// added to the fields of the current logger instance
func (l *Logger) WithFields(fields map[string]string) *Logger {
	newFields := logrus.Fields{}

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
