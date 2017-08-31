package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"

//export kuzzle_wrapper_unset_jwt
func kuzzle_wrapper_unset_jwt() {
	KuzzleInstance.UnsetJwt()
}
