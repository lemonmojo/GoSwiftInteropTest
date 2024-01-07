typedef void* logger_t;

typedef logger_t (*logger_create_func)(void);
typedef void (*logger_destroy_func)(logger_t logger);

typedef void (*logger_log_func)(logger_t logger, const char*);

logger_t logger_create(
	logger_create_func f
) {
    return f();
}

void logger_destroy(
	logger_destroy_func f,
	logger_t logger
) {
    f(logger);
}

void logger_log(
	logger_log_func f,
	logger_t logger,
	const char* message
) {
    f(logger, message);
}