package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import "unsafe"

//export kuzzle_wrapper_logout
func kuzzle_wrapper_logout() *C.char {
	err := KuzzleInstance.Logout()
	if err != nil {
		return C.CString(err.Error())
	}

	return nil
}
