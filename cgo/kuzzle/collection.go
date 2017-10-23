package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
	#include <string.h>
*/
import "C"
import (
	"unsafe"
)

//export kuzzle_wrapper_new_collection
func kuzzle_wrapper_new_collection(k *C.Kuzzle, colName *C.char, index *C.char) *C.collection {
	// TODO Must be freed in C
	col := (*C.collection)(C.calloc(1, C.sizeof_collection))
	// TODO Must be freed in C
	C.strcpy(col.index, index)
	// TODO Must be freed in C
	C.strcpy(col.collection, colName)

	// TODO Must be freed by the Kuzzle destructor
	col.kuzzle = unsafe.Pointer(k)

	return col
}
