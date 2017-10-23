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
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_m_delete_document
func kuzzle_wrapper_collection_m_delete_document(c *C.collection , ids **C.char, idsCount C.uint, options *C.query_options) *C.string_array_result {
	result := (*C.string_array_result)(C.calloc(1, C.sizeof_string_array_result))
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	gIds := goStrings(ids, idsCount)
	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.MDeleteDocument(gIds, opts)

	if err != nil {
		Set_string_array_result_error(result, err)
		return result
	}

	result.result = (**C.char)(C.calloc(C.size_t(len(res)), C.sizeof_char_ptr))
	result.length = C.ulong(len(res))

	cArray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(result.result))[:len(res):len(res)]

	for i, substring := range res {
		cArray[i] = C.CString(substring)
	}

	return result
}
