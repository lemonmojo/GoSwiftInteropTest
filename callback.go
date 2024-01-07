package main

/*
typedef void (*callbackFunc)(int);
extern void call_callbackFunc(callbackFunc f, int number);
*/
import "C"

//export test_callback
func test_callback(number C.int, cb C.callbackFunc) {
	logger := logger_create()
	logger_log(logger, "Hello from Logger")
	logger_destroy(logger)

	C.call_callbackFunc(cb, number)
}
