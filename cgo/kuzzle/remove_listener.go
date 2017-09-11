package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_remove_listener
func kuzzle_wrapper_remove_listener(event C.int) {
	KuzzleInstance.RemoveListener(int(event))
}
