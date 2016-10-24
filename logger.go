package log

import (
	"io"
	"strings"

	"github.com/fatih/color"
)

type LoggerFunc func(m *Message) error
type MarshalFunc func(interface{}) ([]byte, error)

type Logger struct {
	Stdout       io.Writer
	Stderr       io.Writer
	Level        Level
	DefaultLevel Level
	Color        bool
	Marshal      MarshalFunc
	LogFunc      LoggerFunc
}

func (l *Logger) Log(m *Message) error {
	if m.Level > l.Level {
		return nil
	}
	if l.LogFunc != nil {
		return l.LogFunc(m)
	}
	return l.DefaultLogFunc(m)
}

func (l *Logger) DefaultLogFunc(m *Message) error {
	var bs []byte
	if l.Marshal != nil {
		var err error
		_m := *m
		if strings.HasSuffix(m.Text, "\n") {
			_m.Text = m.Text[:len(m.Text)-1]
		}
		bs, err = l.Marshal(_m)
		if err != nil {
			return err
		}
		bs = append(bs, '\n')
	} else {
		if l.Color {
			switch m.Level {
			case LevelCritical, LevelFatal:
				bs = []byte(color.RedString(m.String()))
			case LevelError:
				bs = []byte(color.RedString(m.String()))
			case LevelWarning:
				bs = []byte(color.YellowString(m.String()))
			case LevelInfo:
				bs = []byte(color.GreenString(m.String()))
			case LevelDebug:
				bs = []byte(m.String())
			default:
				bs = []byte(m.String())
			}
		} else {
			bs = []byte(m.String())
		}
	}
	if m.Level < LevelInfo {
		_, err := l.Stderr.Write(bs)
		if err != nil {
			return err
		}
	} else {
		_, err := l.Stdout.Write(bs)
		if err != nil {
			return err
		}
	}
	return nil
}

// Write always returns 0, nil
func (l *Logger) Write(bs []byte) (int, error) {
	l.Log(NewMessage(l.DefaultLevel, string(bs)))
	return 0, nil
}
