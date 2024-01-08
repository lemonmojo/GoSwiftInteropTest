package main

/*
typedef void* logger_t;

typedef logger_t (*logger_create_func)(void);
typedef void (*logger_destroy_func)(logger_t logger);

typedef void (*logger_log_func)(logger_t logger, const char*);

static logger_t logger_create(
	logger_create_func f
) {
    return f();
}

static void logger_destroy(
	logger_destroy_func f,
	logger_t logger
) {
    f(logger);
}

static void logger_log(
	logger_log_func f,
	logger_t logger,
	const char* message
) {
    f(logger, message);
}
*/
import "C"

var _loggerCreateFunc C.logger_create_func
var _loggerDestroyFunc C.logger_destroy_func
var _loggerLogFunc C.logger_log_func

//export logger_configure
func logger_configure(create C.logger_create_func,
	destroy C.logger_destroy_func,
	log C.logger_log_func) {
	_loggerCreateFunc = create
	_loggerDestroyFunc = destroy
	_loggerLogFunc = log
}

type Logger struct {
	logger C.logger_t
}

func (l Logger) Log(message string) {
	C.logger_log(_loggerLogFunc, l.logger, C.CString(message))
}

func (l Logger) Close() {
	C.logger_destroy(_loggerDestroyFunc, l.logger)
}

func NewLogger() Logger {
	return Logger{
		logger: C.logger_create(_loggerCreateFunc),
	}
}
