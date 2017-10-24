package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_flush_queue
func kuzzle_wrapper_flush_queue(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).FlushQueue()
}
