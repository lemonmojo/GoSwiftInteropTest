package main

/*
typedef void* logger_t;

typedef logger_t (*logger_create_func)(void);
typedef void (*logger_destroy_func)(logger_t logger);

typedef void (*logger_log_func)(logger_t logger, const char*);

extern logger_t logger_create(
	logger_create_func f
);

extern void logger_destroy(
	logger_destroy_func f,
	logger_t logger
);

extern void logger_log(
	logger_log_func f,
	logger_t logger,
	const char* message
);
*/
import "C"

var loggerCreateFunc C.logger_create_func
var loggerDestroyFunc C.logger_destroy_func
var loggerLogFunc C.logger_log_func

//export logger_configure
func logger_configure(create C.logger_create_func,
	destroy C.logger_destroy_func,
	log C.logger_log_func) {
	loggerCreateFunc = create
	loggerDestroyFunc = destroy
	loggerLogFunc = log
}

func logger_create() C.logger_t {
	return C.logger_create(loggerCreateFunc)
}

func logger_destroy(logger C.logger_t) {
	C.logger_destroy(loggerDestroyFunc, logger)
}

func logger_log(logger C.logger_t, message string) {
	C.logger_log(loggerLogFunc, logger, C.CString(message))
}
