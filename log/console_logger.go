package log

import (
	"fmt"
	"time"
)

type ConsoleLogger struct {
}

var (
	levelStrMap = map[Level]string{
		LevelTrace: "TRACE",
		LevelDebug: "DEBUG",
		LevelInfo:  " INFO",
		LevelWarn:  " WARN",
		LevelError: "ERROR",
		LevelFatal: "FATAL",
	}

	levelColorPrefixMap = map[Level]string{
		LevelTrace: "\x1b[0;37m",
		LevelDebug: "\x1b[0;36m",
		LevelInfo:  "\x1b[0;34m",
		LevelWarn:  "\x1b[0;33m",
		LevelError: "\x1b[0;31m",
		LevelFatal: "\x1b[0;31;47m",
	}
	colorSuffix = "\x1b[0m"
)

func NewConsoleLogger() ConsoleLogger {
	return ConsoleLogger{}
}

func (c ConsoleLogger) Do(caller *CallerInfo, format string, val ...interface{}) {
	f := fmt.Sprintf("%s[%s][%s][%s:%d] %s%s\n",
		levelColorPrefixMap[caller.Level],
		levelStrMap[caller.Level],
		time.Now().Format("2006/01/02 15:04:05"),
		caller.FilePath,
		caller.LineNumber,
		format, colorSuffix)

	fmt.Printf(f, val...)
}
