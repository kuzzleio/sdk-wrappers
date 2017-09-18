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
  "encoding/json"
)

//export kuzzle_wrapper_collection_get_specifications
func kuzzle_wrapper_collection_get_specifications(c *C.collection, result *C.kuzzle_response, options *C.query_options) C.int {
  var opts types.QueryOptions
  if options != nil {
    opts = SetOptions(options)
  }

  res, err := (*collection.Collection)(c.instance).GetSpecifications(opts)
  if err != nil {
    result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
    return 0
  }

  var jsonRes *C.json_object
  r, _ := json.Marshal(res)

  jsonRes = C.json_tokener_parse(C.CString(string(r)))
  result.result = jsonRes

  return 0
}
