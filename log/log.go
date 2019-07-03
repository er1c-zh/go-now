package log

import "sync"

var (
	logger logs

	once sync.Once
)

func initLogger() {
	once.Do(initLoggerDefault)
}

func initLoggerDefault() {
	logger = NewConsoleLogger()
}

func Trace(format string, val ...interface{}) {
	do(NewCallerInfo(LevelTrace, DefaultDepth), format, val...)
}
func Debug(format string, val ...interface{}) {
	do(NewCallerInfo(LevelDebug, DefaultDepth), format, val...)
}
func Info(format string, val ...interface{}) {
	do(NewCallerInfo(LevelInfo, DefaultDepth), format, val...)
}
func Warn(format string, val ...interface{}) {
	do(NewCallerInfo(LevelWarn, DefaultDepth), format, val...)
}
func Error(format string, val ...interface{}) {
	do(NewCallerInfo(LevelError, DefaultDepth), format, val...)
}
func Fatal(format string, val ...interface{}) {
	do(NewCallerInfo(LevelFatal, DefaultDepth), format, val...)
}

func do(caller *CallerInfo, format string, val ...interface{}) {
	if logger == nil {
		initLogger()
	}
	logger.Do(caller, format, val...)
}
