package lg

import (
	"net/http"
	"net/http/httputil"
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

func (l *Logger) Response(r *http.Response) {
	b, _ := httputil.DumpResponse(r, true)

	l.Debug(string(b))
}

func (l *Logger) Check(err error, args ...interface{}) {
	if err != nil {
		l.Fatal(err.Error(), args...)
	}
}
