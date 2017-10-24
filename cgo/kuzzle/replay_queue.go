package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import "github.com/kuzzleio/sdk-go/kuzzle"

//export kuzzle_wrapper_replay_queue
func kuzzle_wrapper_replay_queue(k *C.Kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).ReplayQueue()
}
