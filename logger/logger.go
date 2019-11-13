package logger

import (
	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
}

// Entry creates a log entry for the pkg and function.
func Entry(pkg string, fn string) *log.Entry {
	return log.
		WithField("pkg", pkg).
		WithField("fn", fn)
}

// For creates a logger for a specific package, function and userID.
func For(pkg string, fn string, userID string) *log.Entry {
	return Entry(pkg, fn).WithField("userID", userID)
}
