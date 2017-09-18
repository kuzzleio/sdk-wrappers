package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
  "unsafe"
  "github.com/kuzzleio/sdk-go/types"
  "github.com/kuzzleio/sdk-go/collection"
)

//export kuzzle_wrapper_collection_count
func kuzzle_wrapper_collection_count(c *C.collection, result *C.int_response, filters *C.filters, options *C.query_options) C.int {
  var opts types.QueryOptions
  if options != nil {
    opts = SetOptions(options)
  }

  res, err := (*collection.Collection)(c.instance).Count(filters, opts)
  if err != nil {
    result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
    return 0
  }

  result.result = C.int(res)

  return 0
}
