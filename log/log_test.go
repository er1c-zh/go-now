package log

import "testing"

func TestBase(t *testing.T) {
	t.Logf("Trace: %d, Debug: %d, Info: %d, Warn: %d, Error: %d, Fatal: %d",
		LevelTrace, LevelDebug, LevelInfo, LevelWarn, LevelError, LevelFatal)

	Trace("num: %d, string: %s, struct: %+v", LevelTrace, "trace", NewCallerInfo(LevelTrace, 1))
	Debug("num: %d, string: %s, struct: %+v", LevelDebug, "debug", NewCallerInfo(LevelDebug, 1))
	Info("num: %d, string: %s, struct: %+v", LevelInfo, "info", NewCallerInfo(LevelInfo, 1))
	Warn("num: %d, string: %s, struct: %+v", LevelWarn, "Warn", NewCallerInfo(LevelWarn, 1))
	Error("num: %d, string: %s, struct: %+v", LevelError, "Error", NewCallerInfo(LevelError, 1))
	Fatal("num: %d, string: %s, struct: %+v", LevelFatal, "Fatal", NewCallerInfo(LevelFatal, 1))
}

func TestCallerInfo(t *testing.T) {
	caller := NewCallerInfo(LevelTrace, 1)

	t.Logf("%+v", caller)
}
