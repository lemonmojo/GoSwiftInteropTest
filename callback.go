package main

/*
typedef void (*callbackFunc) (int number);

static void call_callbackFunc(callbackFunc f, int number) {
    f(number);
}
*/
import "C"

// public to c
//
//export test_callback
func test_callback(number C.int, cb C.callbackFunc) {
	C.call_callbackFunc(cb, number)
}
