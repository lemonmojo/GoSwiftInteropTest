package main

import (
	"C"
	"fmt"
	"strings"
)

// public
func Uppercase(input string) string {
	return strings.ToUpper(input)
}

// public to c
//
//export uppercase_string
func uppercase_string(input *C.char) *C.char {
	logger := NewLogger()
	defer logger.Close()

	logger.Log("Hello from uppercase_string in golang! 🥳")

	inputGo := C.GoString(input)
	outputGo := Uppercase(inputGo)
	outputC := C.CString(outputGo)

	logger.Log(fmt.Sprintf("uppercase_string Input: \"%s\"; Output: \"%s\"", inputGo, outputGo))

	return outputC
}
