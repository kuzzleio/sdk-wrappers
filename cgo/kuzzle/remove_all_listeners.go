package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_remove_all_listeners
func kuzzle_wrapper_remove_all_listeners(event C.int) {
	KuzzleInstance.RemoveAllListeners(int(event))
}
