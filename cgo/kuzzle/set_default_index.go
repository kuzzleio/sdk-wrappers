package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_set_default_index
func kuzzle_wrapper_set_default_index(index *C.char) int {
	err := KuzzleInstance.SetDefaultIndex(C.GoString(index))
	if err != nil {
		return C.int(C.EINVAL)
	}

	return 0
}
