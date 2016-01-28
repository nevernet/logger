package logger

import "github.com/astaxie/beego/logs"

var Log *logs.BeeLogger

func init() {
	Log = logs.NewLogger(10000)
	Log.SetLogger("console", `{"level": 10}`)
	Log.SetLevel(logs.LevelDebug)
	Log.EnableFuncCallDepth(true)
	Log.SetLogFuncCallDepth(3)
	Log.Async()
}

func Trace(format string, v ...interface{}) {
	Log.Trace(format, v...)
}

func Debug(format string, v ...interface{}) {
	Log.Debug(format, v...)
}

func Info(format string, v ...interface{}) {
	Log.Info(format, v...)
}

func Warn(format string, v ...interface{}) {
	Log.Warn(format, v...)
}

func Error(format string, v ...interface{}) {
	Log.Error(format, v...)
}

func Critical(format string, v ...interface{}) {
	Log.Critical(format, v...)
}

func Flush() {
	Log.Flush()
}

func Close() {
	Log.Close()
}
