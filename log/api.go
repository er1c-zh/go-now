package log

type Logger interface {
	Do(caller *CallerInfo, format string, val ...interface{})
	Flush()
}
