// Copyright 2014 Daniel Qin. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logger provide a general log interface
// Usage:
//
// import "github.com/nevernet/logger"
//	logger.Trace("trace")
//	logger.Info("info")
//	logger.Warn("warning")
//	logger.Debug("debug")
//	logger.Critical("critical")
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

// SetLogger reset the logger, default is console logger
func SetLogger(adapterName string, configs ...string) {
	log.SetLogger(adapterName, configs...)
}

// GetLogger return the logger
func GetLogger() *logs.BeeLogger {
	return log
}

// SetFileLogger is helper function for set file adapter
func SetFileLogger(filename string, maxlines int, maxsize int, daily bool, maxdays int, rotate bool, level string, separate string) {
	configs := fmt.Sprintf(`{"filename":"%s", "maxlines":%d, "maxsize":%d, "daily":%t, "maxdays":%d, "rotate":%t, "level":"%s", "separate":"%s"}`,
		filename, maxlines, maxsize, daily, maxdays, rotate, level, separate)
	log.SetLogger(logs.AdapterFile, configs)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)
	log.Async()
}

// Trace log
func Trace(format string, v ...interface{}) {
	log.Trace(format, v...)
}

// Debug log
func Debug(format string, v ...interface{}) {
	log.Debug(format, v...)
}

// Info log
func Info(format string, v ...interface{}) {
	log.Info(format, v...)
}

// Warn log
func Warn(format string, v ...interface{}) {
	log.Warn(format, v...)
}

// Error log
func Error(format string, v ...interface{}) {
	log.Error(format, v...)
}

// Critical log
func Critical(format string, v ...interface{}) {
	log.Critical(format, v...)
}

// Flush log
func Flush() {
	log.Flush()
}

// Close log
func Close() {
	log.Close()
}
