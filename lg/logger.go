package lg

import (
	"time"

	"github.com/mgutz/logxi/v1"
)

type Logger struct {
	log.Logger
}

func New(name string) *Logger {
	return &Logger{
		Logger: log.New(name),
	}
}

func (l *Logger) Start() *Stopwatch {
	return &Stopwatch{
		start:  time.Now(),
		logger: l,
	}
}

func (l *Logger) Check(err error, args ...interface{}) {
	if err != nil {
		l.Fatal(err.Error(), args...)
	}
}
