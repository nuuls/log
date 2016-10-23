package log

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type caller struct {
	file string
	fn   string
	pkg  string
	line int
}

func Caller(skip int) caller {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		return caller{}
	}
	p := filepath.ToSlash(file)
	pkg := p
	spl := strings.Split(p, "/")
	if len(spl) > 1 {
		pkg = spl[len(spl)-2]
	}
	return caller{
		pkg:  pkg,
		file: filepath.Base(file),
		line: line,
	}
}

func (c caller) String() string {
	return c.pkg + "/" + c.file + ":" + strconv.Itoa(c.line)
}
