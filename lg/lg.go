package lg

import (
	"github.com/mgutz/logxi/v1"
)

func Fatal(msg string, args ...interface{}) {
	log.Fatal(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	log.Warn(msg, args...)
}

func Info(msg string, args ...interface{}) {
	log.Info(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	log.Debug(msg, args...)
}

func Check(err error, args ...interface{}) {
	if err != nil {
		log.Fatal(err.Error(), args...)
	}
}

type F map[string]interface{}

func (f F) Args() []interface{} {
	var args []interface{}

	for k, v := range f {
		args = append(args, k, v)
	}

	return args
}
