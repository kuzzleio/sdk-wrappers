package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_disconnect
func kuzzle_wrapper_disconnect(k *C.Kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).Disconnect()
}
