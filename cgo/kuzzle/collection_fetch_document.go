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

//export kuzzle_wrapper_collection_fetch_document
func kuzzle_wrapper_collection_fetch_document(c *C.collection, result *C.document, id *C.char, options *C.query_options) C.int {
  var opts types.QueryOptions
  if options != nil {
    opts = SetQueryOptions(options)
  }

  res, err := (*collection.Collection)(c.instance).FetchDocument(C.GoString(id), opts)
  if err != nil {
    if err.Error() == "Collection.FetchDocument: document id required" {
      return C.int(C.EINVAL)
    } else {
      result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
      return 0
    }
  }

  result.instance = unsafe.Pointer(&res)

  return 0
}
