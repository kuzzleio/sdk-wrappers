package main
/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_replay_queue
func kuzzle_wrapper_replay_queue() {
	KuzzleInstance.ReplayQueue()
}