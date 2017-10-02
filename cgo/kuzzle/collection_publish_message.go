package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
/* TODO
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_collection_publish_message
func kuzzle_wrapper_collection_publish_message(c *C.collection, result *C.bool_result, message *C.json_object, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	jp := JsonParser{}
	jp.Parse(message)

	res, err := (*collection.Collection)(c.instance).PublishMessage(jp.GetContent(), opts)
	if err != nil {
		result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
		return
	}

	var r uint

	if res.Published {
		r = 1
	}

	result.result = r

	return
}
*/