package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_logout
func kuzzle_wrapper_logout(k *C.kuzzle) *C.char {
	err := (*kuzzle.Kuzzle)(k.instance).Logout()
	if err != nil {
		// TODO Must be freed in C
		return C.CString(err.Error())
	}

	return nil
}
