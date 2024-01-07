package main

import (
	"C"
	"strings"
)

//export uppercase_string
func uppercase_string(input *C.char) *C.char {
	logger := NewLogger()
	defer logger.Close()

	logger.Log("Hello from uppercase_string in golang! 🥳")

	inputGo := C.GoString(input)
	upperGo := strings.ToUpper(inputGo)
	upperC := C.CString(upperGo)

	return upperC
}

// We need an entry point; it's ok for this to be empty
func main() {}
