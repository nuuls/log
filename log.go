package log

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Import can be used to prevent goimports from importing "log"
// var _ = log.Import()
func Import() struct{} {
	return struct{}{}
}

// Level xd
type Level int

func (l Level) String() string {
	switch l {
	case LevelFatal:
		return "FATAL"
	case LevelCritical:
		return "CRIT"
	case LevelError:
		return "ERRO"
	case LevelWarning:
		return "WARN"
	case LevelInfo:
		return "INFO"
	case LevelDebug:
		return "DEBU"
	}
	return "undefined"
}

// Available Levels
const (
	LevelFatal Level = iota
	LevelCritical
	LevelError
	LevelWarning
	LevelInfo
	LevelDebug
)

var DefaultLogger = &Logger{
	Stdout:       os.Stdout,
	Stderr:       os.Stderr,
	Level:        LevelDebug,
	DefaultLevel: LevelDebug,
	Color:        true,
}

var (
	// TimeFormat is the Default Time Format
	TimeFormat = "2006-01-02 15:04:05.000"

	Stack Level = LevelError

	Loggers = []*Logger{}

	CallerStrLen = 35
)

func AddLogger(l *Logger) {
	if l.Stderr == nil {
		if l.Stdout == nil {
			panic("no writer")
		}
		l.Stderr = l.Stdout
	}
	if l.Stdout == nil {
		l.Stdout = l.Stderr
	}
	Loggers = append(Loggers, l)
}

type Message struct {
	Level       Level     `json:"level"`
	LevelString string    `json:"levelString"`
	Time        time.Time `json:"time"`
	Caller      string    `json:"caller"`
	Stack       string    `json:"stack,omitempty"`
	Text        string    `json:"text"`
	args        []interface{}
}

func NewMessage(level Level, a ...interface{}) *Message {
	m := &Message{
		Level:       level,
		LevelString: level.String(),
		Time:        time.Now(),
		Caller:      Caller(2).String(),
		args:        a,
		Text:        fmt.Sprintln(a...),
	}
	return m
}

func (m *Message) String() string {
	cLen := CallerStrLen - len(m.Caller)
	if cLen < 1 {
		cLen = 1
	}
	_caller := m.Caller + strings.Repeat(" ", cLen)
	return fmt.Sprintf("%s %s %s %s",
		m.Time.Format(TimeFormat),
		_caller,
		m.Level,
		m.Text)
}

func Debug(a ...interface{}) {
	m := NewMessage(LevelDebug, a...)
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Debugf(format string, a ...interface{}) {
	m := NewMessage(LevelDebug, fmt.Sprintf(format, a...))
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Info(a ...interface{}) {
	m := NewMessage(LevelInfo, a...)
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Infof(format string, a ...interface{}) {
	m := NewMessage(LevelInfo, fmt.Sprintf(format, a...))
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Warning(a ...interface{}) {
	m := NewMessage(LevelWarning, a...)
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Warningf(format string, a ...interface{}) {
	m := NewMessage(LevelWarning, fmt.Sprintf(format, a...))
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Error(a ...interface{}) {
	m := NewMessage(LevelError, a...)
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Errorf(format string, a ...interface{}) {
	m := NewMessage(LevelError, fmt.Sprintf(format, a...))
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Critical(a ...interface{}) {
	m := NewMessage(LevelCritical, a...)
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Criticalf(format string, a ...interface{}) {
	m := NewMessage(LevelCritical, fmt.Sprintf(format, a...))
	for _, l := range Loggers {
		l.Log(m)
	}
}

func Fatal(a ...interface{}) {
	m := NewMessage(LevelFatal, a...)
	for _, l := range Loggers {
		l.Log(m)
	}
	os.Exit(1)
}

func Fatalf(format string, a ...interface{}) {
	m := NewMessage(LevelFatal, fmt.Sprintf(format, a...))
	for _, l := range Loggers {
		l.Log(m)
	}
}
