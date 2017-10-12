package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
)

//export kuzzle_wrapper_collection_publish_message
func kuzzle_wrapper_collection_publish_message(c *C.collection, result *C.bool_result, message *C.json_object, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	res, err := (*collection.Collection)(c.instance).PublishMessage(JsonCConvert(message).(map[string]interface{}), opts)
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