package log

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileLogger struct {
	ConsoleLogger ConsoleLogger
	path          string
	file          *os.File
}

func NewFileLogger(path string) FileLogger {
	l := FileLogger{
		ConsoleLogger: NewConsoleLogger(),
		path:          path,
	}
	if err := os.MkdirAll(filepath.Dir(path), os.FileMode(0755)); err != nil {
		fmt.Printf("MkdirAll(%s) fail: %s\n", path, err.Error())
		return l
	}
	file, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, os.FileMode(0644))
	if err != nil {
		fmt.Printf("MkdirAll(%s) fail: %s\n", path, err.Error())
		return l
	}
	l.file = file
	return l
}

func (f FileLogger) Do(caller *CallerInfo, format string, val ...interface{}) {
	if f.file == nil {
		f.ConsoleLogger.Do(caller, format, val...)
		return
	}
	s := fmt.Sprintf("[%s][%s][%s:%d] %s\n",
		levelStrMap[caller.Level],
		time.Now().Format("2006/01/02 15:04:05"),
		caller.FilePath,
		caller.LineNumber,
		format)
	_, _ = f.file.WriteString(fmt.Sprintf(s, val...))
}

func (f FileLogger) Flush() {
	if f.file == nil {
		f.ConsoleLogger.Flush()
		return
	}
	_ = f.file.Sync()
}
