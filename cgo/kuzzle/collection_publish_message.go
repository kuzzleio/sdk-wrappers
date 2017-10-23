package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_publish_message
// TODO
func kuzzle_wrapper_collection_publish_message(c *C.collection, result *C.bool_result, message *C.json_object, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.PublishMessage(JsonCConvert(message).(map[string]interface{}), opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	var r C.uint

	if res.Published {
		r = 1
	} else {
		r = 0
	}

	result.result = r

	return
}