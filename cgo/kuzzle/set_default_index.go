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

//export kuzzle_wrapper_set_default_index
func kuzzle_wrapper_set_default_index(k *C.kuzzle, index *C.char) C.int {
	err := (*kuzzle.Kuzzle)(k.instance).SetDefaultIndex(C.GoString(index))
	if err != nil {
		return C.int(C.EINVAL)
	}

	return 0
}
