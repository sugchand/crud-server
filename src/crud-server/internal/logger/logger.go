// Copyright 2019 Sugesh Chandran
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package logger is a module used across the application to manage logs.
// There will be multiple instances of this logger present in the application
package logger

import (
	"fmt"
	"io"
	"log"
	"os"
)

// LogLeveltype :- Global type to define the log level.
type LogLeveltype uint64

// Logging is a wrapper struct to track the module logging
// Log.logger is thread safe by default.
// It can be invoked from different threads
// at the same time without any issues.
type Logging struct {
	// loglevel can be Trace/Info/Warning/Error
	currloglevel LogLeveltype
	//format of the logs to be printed.
	logformatFlags int
	tracerLogger   *log.Logger
	infoLogger     *log.Logger
	warningLogger  *log.Logger
	errorLogger    *log.Logger
	fp             io.Writer
}

// Log level constants.
const (
	Trace = iota + 1
	Info
	Warning
	Error
)

// LogInit function to initilized the Logging instance.
// Application can have only single Logging instance to keep limited memory
// usage.
func (logger *Logging) LogInit(loglevel LogLeveltype,
	filepath string) {
	var err error
	var stdoutHandler io.Writer
	//var fp *os.File
	stdoutHandler = os.Stdout
	logger.currloglevel = loglevel
	logger.logformatFlags = log.Ldate | log.Ltime
	if len(filepath) == 0 {
		logger.fp = stdoutHandler
	} else {
		logger.fp, err = os.OpenFile(filepath,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			//Failed to create/delete the file, so just use standard output
			fmt.Printf("Cannot use file for logging, defaulting to stdout"+
				" err : %s", err)
			logger.fp = stdoutHandler
		}
	}
	logger.initloggers()

}

func (logger *Logging) initloggers() {
	logger.tracerLogger = log.New(logger.fp, "TRACE: ", logger.logformatFlags)
	logger.infoLogger = log.New(logger.fp, "INFO: ", logger.logformatFlags)
	logger.warningLogger = log.New(logger.fp, "WARNING: ", logger.logformatFlags)
	logger.errorLogger = log.New(logger.fp, "ERROR: ", logger.logformatFlags)
}

// Trace :- log to trace the call routine.
func (logger *Logging) Trace(msgfmt string, args ...interface{}) {
	if logger.currloglevel > Trace {
		return
	}
	logger.tracerLogger.Printf(msgfmt, args...)
}

// Info :- log function to write info
func (logger *Logging) Info(msgfmt string, args ...interface{}) {
	if logger.currloglevel > Info {
		return
	}
	logger.infoLogger.Printf(msgfmt, args...)
}

// Warning :- Function to track the warning logs
func (logger *Logging) Warning(msgfmt string, args ...interface{}) {
	if logger.currloglevel > Warning {
		return
	}
	logger.warningLogger.Printf(msgfmt, args...)
}

func (logger *Logging) Error(msgfmt string, args ...interface{}) {
	if logger.currloglevel > Error {
		return
	}
	logger.errorLogger.Printf(msgfmt, args...)
}
