package log

import (
	"bytes"
	"fmt"
	"sync"
	"time"
)

type ConsoleLogger struct {
	chLog   chan string
	chFlush chan struct{}
}

const maxBufLength = 32
const maxFlushInterval = 200 * time.Millisecond

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

	consoleLoggerOnce sync.Once
)

func NewConsoleLogger() ConsoleLogger {
	l := ConsoleLogger{
		chLog:   make(chan string),
		chFlush: make(chan struct{}),
	}
	consoleLoggerOnce.Do(func() {
		go l.flushBackground()
	})
	return l
}

func (c ConsoleLogger) Do(caller *CallerInfo, format string, val ...interface{}) {
	f := fmt.Sprintf("%s[%s][%s][%s:%d] %s%s\n",
		levelColorPrefixMap[caller.Level],
		levelStrMap[caller.Level],
		time.Now().Format("2006/01/02 15:04:05"),
		caller.FilePath,
		caller.LineNumber,
		format, colorSuffix)

	c.chLog <- fmt.Sprintf(f, val...)
}

func (c ConsoleLogger) Flush() {
	c.chFlush <- struct{}{}
	<-c.chFlush
}

func (c ConsoleLogger) flushBackground() {
	ticker := time.NewTicker(maxFlushInterval)

	count := 0
	buf := bytes.NewBufferString("")

	flush := func() {
		fmt.Print(buf.String())
		buf.Reset()
		count = 0
	}

	for {
		select {
		case s, ok := <-c.chLog:
			if !ok {
				continue
			}
			buf.WriteString(s)
			count++
			if count >= maxBufLength {
				flush()
			}
		case <-c.chFlush:
			if count > 0 {
				flush()
			}
			c.chFlush <- struct{}{}
		case <-ticker.C:
			if count > 0 {
				flush()
			}
		}
	}
}
