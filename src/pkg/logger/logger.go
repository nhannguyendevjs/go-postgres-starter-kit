package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// Log is a global logger instance
var Log = logrus.New()

func init() {
	// Set log output to stdout
	Log.Out = os.Stdout

	// Set log format
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: time.RFC3339,
	})

	// Set log level (can be adjusted based on environment)
	Log.SetLevel(logrus.InfoLevel)
}

// Info logs an info message
func Info(message string, fields logrus.Fields) {
	Log.WithFields(fields).Info(message)
}

// Error logs an error message
func Error(message string, fields logrus.Fields) {
	Log.WithFields(fields).Error(message)
}

// Debug logs a debug message
func Debug(message string, fields logrus.Fields) {
	Log.WithFields(fields).Debug(message)
}
