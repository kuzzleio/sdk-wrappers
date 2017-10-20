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

//export kuzzle_wrapper_collection_create
func kuzzle_wrapper_collection_create(c *C.collection, options *C.query_options) *C.ack_result {
	result := (*C.ack_result)(C.calloc(1, C.sizeof_ack_result))

	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.Create(opts)

	if err != nil {
		Set_ack_result_error(result, err)
		return result
	}

	if res.Acknowledged {
		result.acknowledged = 1
	} else {
		result.acknowledged = 0
	}

	if res.ShardsAcknowledged {
		result.shardsAcknowledged = 1
	} else {
		result.shardsAcknowledged = 0
	}

	return result
}
