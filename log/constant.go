package log

type Level int8

const (
	LevelTrace Level = 1 << (iota)
	LevelDebug
	LevelInfo
	LevelWarn
	LevelError
	LevelFatal
)
