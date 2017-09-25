package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_collection_m_delete_document
func kuzzle_wrapper_collection_m_delete_document(c *C.collection, result *C.string_array_result, ids **C.char, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	gIds := goStrings(ids)

	res, err := (*collection.Collection)(c.instance).MDeleteDocument(gIds, opts)
	if err != nil {
		if err.Error() == "Collection.MDeleteDocument: please provide at least one id of document to delete" {
			return C.int(C.EINVAL)
		} else {
			result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
			return 0
		}
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
	return 0
}
