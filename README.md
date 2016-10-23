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
		log.Fatal(err)
	}
	jsonLogger := &log.Logger{
		Level:        log.LevelError,
		Stdout:       file,
		Stderr:       file,
		DefaultLevel: log.LevelError,
		MarshalFunc:  json.Marshal,
	}
	log.AddLogger(jsonLogger)
	log.Debug("test 123")
	log.Infof("%s 123", "Kappa")
	log.Error("error")
}

  ```
