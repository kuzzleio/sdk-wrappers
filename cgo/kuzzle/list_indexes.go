package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
*/
import "C"
import (
	"unsafe"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_list_indexes
func kuzzle_wrapper_list_indexes(result *C.string_array_result, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetOptions(options)
	}

	res, err := KuzzleInstance.ListIndexes(opts)
	if err != nil {
		(*result).error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
	}

	cArray := C.malloc(C.size_t(len(res)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	a := (*[1<<30 - 1]*C.char)(cArray)

	idx := 0
	for _, substring := range res {
		a[idx] = C.CString(substring)
		idx += 1
	}
	a[idx] = nil

	(*result).result = (**C.char)(cArray)
}
