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
func kuzzle_wrapper_collection_create(c *C.collection, result *C.ack_response, options *C.query_options) {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	res, err := col.Create(opts)

	if err != nil {
		result.error = ToCString_2048(err.Error())
		return
	}

	var ack, shardsAck C.uint

	if res.Acknowledged {
		ack = 1
	} else {
		ack = 0
	}

	if res.ShardsAcknowledged {
		shardsAck = 1
	} else {
		shardsAck = 0
	}

	*result = C.ack_response{
		acknowledged:       ack,
		shardsAcknowledged: shardsAck,
	}
}
