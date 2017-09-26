package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import "github.com/kuzzleio/sdk-go/kuzzle"

//export kuzzle_wrapper_start_queuing
func kuzzle_wrapper_start_queuing(k *C.Kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).StartQueuing()
}
