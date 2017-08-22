package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <json/json.h>
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_disconnect
func kuzzle_wrapper_disconnect() {
	KuzzleInstance.Disconnect()
}