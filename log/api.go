package log

type logs interface {
	Do(caller *CallerInfo, format string, val ...interface{})
	Flush()
}
