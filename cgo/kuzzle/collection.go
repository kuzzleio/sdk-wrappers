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
func Collection(c *C.collection, k *C.kuzzle, colName string, index string) *C.collection {
	instance := collection.NewCollection((*kuzzle.Kuzzle)(k.instance), colName, index)

	c.instance = unsafe.Pointer(instance)

	return c
}
