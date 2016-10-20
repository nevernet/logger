package logger

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

type logBridge struct {
}

var log *logs.BeeLogger

func init() {
	log = logs.NewLogger(10000)
	log.SetLogger("console", `{"level": 10}`)
	log.SetLevel(logs.LevelDebug)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)
	log.Async()
}

func SetLogger(adapterName string, configs ...string) {
	log.SetLogger(adapterName, configs...)
}

func SetFileLogger(filename string, maxlines int, maxsize int, daily bool, maxdays int, rotate bool, level string, separate string) {
	configs := fmt.Sprintf(`{"filename":"%s", "maxlines":%d, "maxsize":%d, "daily":%t, "maxdays":%d, "rotate":%t, "level":"%s", "separate":"%s"}`,
		filename, maxlines, maxsize, daily, maxdays, rotate, level, separate)
	log.SetLogger(logs.AdapterFile, configs)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)
	log.Async()
}

func Trace(format string, v ...interface{}) {
	log.Trace(format, v...)
}

func Debug(format string, v ...interface{}) {
	log.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	log.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	log.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	log.Error(format, v...)
}

func Critical(format string, v ...interface{}) {
	log.Critical(format, v...)
}

func Flush() {
	log.Flush()
}

func Close() {
	log.Close()
}
