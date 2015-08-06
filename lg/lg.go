package lg

import "net/http"

var DefaultLog *Logger

func init() {
	DefaultLog = New("~")
}

func Fatal(msg string, args ...interface{}) {
	DefaultLog.Fatal(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	DefaultLog.Warn(msg, args...)
}

func Info(msg string, args ...interface{}) {
	DefaultLog.Info(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	DefaultLog.Debug(msg, args...)
}

func Check(err error, args ...interface{}) {
	if err != nil {
		DefaultLog.Fatal(err.Error(), args...)
	}
}

func Response(r *http.Response) {
	DefaultLog.Response(r)
}

type F map[string]interface{}

func (f F) Args() []interface{} {
	var args []interface{}

	for k, v := range f {
		args = append(args, k, v)
	}

	return args
}
