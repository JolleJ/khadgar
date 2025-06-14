package test

import (
	"jollej/db-scout/lib/prettylog"
	"testing"
)

func TestInfoLog(t *testing.T) {

	log := prettylog.NewPrettyLog()

	log.Info("This is an info log test.")

	if log == nil {
		t.Error("Info log test failed.")
	}
}

func TestDebugLog(t *testing.T) {
	log := prettylog.NewPrettyLog()

	log.Debug("This is a debug log test.")

	if log == nil {
		t.Error("Debug log test failed.")
	}
}

func TestErrorLog(t *testing.T) {
	log := prettylog.NewPrettyLog()

	log.Error("This is an error log test.")

	if log == nil {
		t.Error("Error log test failed.")
	}
}
