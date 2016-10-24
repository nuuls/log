# log
go logging library

[![godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/nuuls/log)
```go
package main

import (
	"encoding/json"
	"os"

	"github.com/nuuls/log"
)

func main() {
	logger := &log.Logger{
		Level:        log.LevelDebug,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
		DefaultLevel: log.LevelDebug,
		Color:        true,
	}
	log.AddLogger(logger)

	file, err := os.OpenFile("test.log", os.O_CREATE|os.O_APPEND, 666)
	if err != nil {
		log.Critical(err)
	}
	var jsonLogger *log.Logger
	jsonLogger = &log.Logger{
		Level:        log.LevelError,
		Stdout:       file,
		Stderr:       file,
		DefaultLevel: log.LevelError,
		Marshal:      json.Marshal,
		LogFunc: func(m *log.Message) error {
			err := jsonLogger.DefaultLogFunc(m)
			if err != nil {
				logger.Log(log.NewMessage(log.LevelCritical, "cannot write to file test.log:", err))
			}
			return nil
		},
	}
	log.AddLogger(jsonLogger)
	log.Debug("debug")
	log.Infof("%s infof", "test")
	log.Error("error")
	log.Fatal("fatal")
}
  ```
console output:

<img src="https://i.nuuls.com/a4jE.png">

test.log:
```json
{"level":2,"levelString":"ERRO","time":"2016-10-24T14:18:08.483591+02:00","caller":"test/test.go:42","text":"error"}
{"level":0,"levelString":"FATAL","time":"2016-10-24T14:18:08.4840918+02:00","caller":"test/test.go:43","text":"fatal"}
```
