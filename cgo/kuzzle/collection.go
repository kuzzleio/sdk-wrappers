package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
	#include <string.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"unsafe"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export Collection
func Collection(c *C.collection, k *C.Kuzzle, colName *C.char, index *C.char) *C.collection {
	instance := collection.NewCollection((*kuzzle.Kuzzle)(k.instance), C.GoString(colName), C.GoString(index))

	c.instance = unsafe.Pointer(instance)

	return c
}
