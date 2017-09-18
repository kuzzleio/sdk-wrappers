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

//export kuzzle_wrapper_collection_search
func kuzzle_wrapper_collection_search(c *C.collection, result *C.kuzzle_search_response, search_filters *C.search_filters, options *C.query_options) C.int {
  var opts types.QueryOptions
  if options != nil {
    opts = SetOptions(options)
  }

  // TODO: Initialize Search Filters from C to Go
  filters := types.SearchFilters{}

  res, err := (*collection.Collection)(c.instance).Search(filters, opts)
  if err != nil {
    result.error = *(*[2048]C.char)(unsafe.Pointer(C.CString(err.Error())))
    return 0
  }

  var jsonRes *C.json_object
  r, _ := json.Marshal(res)

  jsonRes = C.json_tokener_parse(C.CString(string(r)))
  result.result.hits = jsonRes
  result.result.total = C.int(res.Total)

  return 0
}
