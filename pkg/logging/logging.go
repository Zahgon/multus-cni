// Copyright (c) 2018 Intel Corporation
// Copyright (c) 2021 Multus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package logging

import (
	"io"
	"time"

	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Level type
type Level uint32

// PanicLevel...MaxLevel indicates the logging level
const (
	PanicLevel Level = iota
	ErrorLevel
	VerboseLevel
	DebugLevel
	MaxLevel
	UnknownLevel
)

var loggingStderr bool
var loggingW io.Writer
var loggingLevel Level
var logger *lumberjack.Logger

const defaultTimestampFormat = time.RFC3339

// LogOptions specifies the configuration of the log
type LogOptions struct {
	MaxAge     *int  `json:"maxAge,omitempty"`
	MaxSize    *int  `json:"maxSize,omitempty"`
	MaxBackups *int  `json:"maxBackups,omitempty"`
	Compress   *bool `json:"compress,omitempty"`
}

// SetLogOptions set the LoggingOptions of NetConf
func SetLogOptions(options *LogOptions) {
	_ = "STUB: not implemented"
	// logger is used only if filname is supplied
	return
}

// give some default value

func (l Level) String() string { _ = "STUB: not implemented"; return "" }

func printf(level Level, format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Debugf prints logging if logging level >= debug
func Debugf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Verbosef prints logging if logging level >= verbose
func Verbosef(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf prints logging if logging level >= error
func Errorf(format string, a ...interface{}) error { _ = "STUB: not implemented"; return nil }

// Panicf prints logging plus stack trace. This should be used only for unrecoverable error
func Panicf(format string, a ...interface{}) { _ = "STUB: not implemented"; return }

// GetLoggingLevel gets current logging level
func GetLoggingLevel() Level { _ = "STUB: not implemented"; return *new(Level) }

func getLoggingLevel(levelStr string) Level { _ = "STUB: not implemented"; return *new(Level) }

// SetLogLevel sets logging level
func SetLogLevel(levelStr string) { _ = "STUB: not implemented"; return }

// SetLogStderr sets flag for logging stderr output
func SetLogStderr(enable bool) { _ = "STUB: not implemented"; return }

// SetLogFile sets logging file
func SetLogFile(filename string) {
	_ = "STUB: not implemented"
	// logger is used only if filname is supplied
	return
}

func init() {
	loggingStderr = true
	loggingW = nil
	loggingLevel = PanicLevel
	logger = nil
}
