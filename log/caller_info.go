package log

import (
	"runtime"
	"strings"
)

const DefaultDepth int = 2

type CallerInfo struct {
	Level      Level
	FilePath   string
	LineNumber int
}

var (
	unknownCallerInfo = &CallerInfo{
		Level:      LevelFatal,
		FilePath:   "unknown",
		LineNumber: -1,
	}
)

func NewCallerInfo(level Level, depth int) *CallerInfo {
	if depth < 0 {
		depth = DefaultDepth
	}
	_, f, ln, ok := runtime.Caller(depth)

	if !ok {
		return unknownCallerInfo
	}

	i := strings.LastIndex(f, "/")

	return &CallerInfo{
		Level:      level,
		FilePath:   f[i+1:],
		LineNumber: ln,
	}
}
