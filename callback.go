package main

/*
typedef void (*callbackFunc)(int);
extern void call_callbackFunc(callbackFunc f, int number);
*/
import "C"

//export test_callback
func test_callback(number C.int, cb C.callbackFunc) {
	logger := NewLogger()
	logger.Log("Hello from Logger! 🥳")
	logger.Close()

	C.call_callbackFunc(cb, number)
}
