package main

/*
typedef void* logger_t;
typedef unsigned int LogLevel_t;

typedef logger_t (*logger_create_func)(void);
typedef void (*logger_destroy_func)(logger_t logger);

typedef void (*logger_log_func)(logger_t logger, const char*, LogLevel_t);

static logger_t logger_create(logger_create_func f) {
    return f();
}

static void logger_destroy(logger_destroy_func f, logger_t logger) {
    f(logger);
}

static void logger_log(logger_log_func f, logger_t logger, const char* message, LogLevel_t logLevel) {
    f(logger, message, logLevel);
}
*/
import "C"

// public
type Logger struct {
	loggerPtr C.logger_t
}

type LogLevel uint32

const (
	LogLevelDebug LogLevel = 0
	LogLevelInfo  LogLevel = 1
	LogLevelWarn  LogLevel = 2
	LogLevelError LogLevel = 3
)

// public
func (l Logger) Log(message string, logLevel LogLevel) {
	C.logger_log(_loggerDelegate.logFunc, l.loggerPtr, C.CString(message), C.uint(logLevel))
}

// public
func (l Logger) Close() {
	C.logger_destroy(_loggerDelegate.destroyFunc, l.loggerPtr)
}

// public
func NewLogger() Logger {
	return Logger{
		loggerPtr: C.logger_create(_loggerDelegate.createFunc),
	}
}

// private
type _LoggerDelegate struct {
	createFunc  C.logger_create_func
	destroyFunc C.logger_destroy_func
	logFunc     C.logger_log_func
}

// private
var _loggerDelegate *_LoggerDelegate

// public to c
//
//export logger_configure
func logger_configure(create C.logger_create_func, destroy C.logger_destroy_func, log C.logger_log_func) {
	del := _LoggerDelegate{
		createFunc:  create,
		destroyFunc: destroy,
		logFunc:     log,
	}

	_loggerDelegate = &del
}
