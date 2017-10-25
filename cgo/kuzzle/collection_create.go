package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
	#include <stdlib.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_create
func kuzzle_wrapper_collection_create(c *C.collection, options *C.query_options) *C.ack_result {
	result := (*C.ack_result)(C.calloc(1, C.sizeof_ack_result))
	opts := SetQueryOptions(options)

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.Create(opts)

	if err != nil {
		Set_ack_result_error(result, err)
		return result
	}

	result.acknowledged = C.bool(res.Acknowledged)
	result.shards_acknowledged = C.bool(res.ShardsAcknowledged)

	return result
}
