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

//export kuzzle_wrapper_collection_get_mapping
func kuzzle_wrapper_collection_get_mapping(c *C.collection, result *C.collection_mapping, options *C.query_options) C.int {
  var opts types.QueryOptions
  if options != nil {
    opts = SetQueryOptions(options)
  }

  res, err := (*collection.Collection)(c.instance).GetMapping(opts)
  if err != nil {
    result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
    return 0
  }

  result.instance = unsafe.Pointer(&res)

  return 0
}
