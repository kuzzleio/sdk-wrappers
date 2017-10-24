package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"
import (
	"unsafe"
)

// Allocates memory
//export kuzzle_wrapper_new_collection
func kuzzle_wrapper_new_collection(k *C.Kuzzle, colName *C.char, index *C.char) *C.collection {
	col := (*C.collection)(C.calloc(1, C.sizeof_collection))
	col.index = C.CString(C.GoString(index))
	col.collection = C.CString(C.GoString(colName))
	col.kuzzle = unsafe.Pointer(k)

	return col
}
