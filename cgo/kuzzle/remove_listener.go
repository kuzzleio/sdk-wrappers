package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import "github.com/kuzzleio/sdk-go/kuzzle"

//export kuzzle_wrapper_remove_listener
func kuzzle_wrapper_remove_listener(k *C.Kuzzle, event C.int) {
	(*kuzzle.Kuzzle)(k.instance).RemoveListener(int(event))
}
