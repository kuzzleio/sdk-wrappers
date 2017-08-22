package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_disconnect
func kuzzle_wrapper_flush_queue() {
	KuzzleInstance.FlushQueue()
}