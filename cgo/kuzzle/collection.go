package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"

// Allocates memory
//export kuzzle_wrapper_new_collection
func kuzzle_wrapper_new_collection(k *C.kuzzle, colName *C.char, index *C.char) *C.collection {
	col := (*C.collection)(C.calloc(1, C.sizeof_collection))
	col.index = C.CString(C.GoString(index))
	col.collection = C.CString(C.GoString(colName))
	col.kuzzle = k

	return col
}

//export kuzzle_wrapper_collection_create
func kuzzle_wrapper_collection_create(c *C.collection, options *C.query_options) *C.ack_result {
	res, err := cToGoCollection(c).Create(SetQueryOptions(options))
	return goToCAckResult(res, err)
}

//export kuzzle_wrapper_collection_publish_message
func kuzzle_wrapper_collection_publish_message(c *C.collection, message *C.json_object, options *C.query_options) *C.bool_result {
	res, err := cToGoCollection(c).PublishMessage(JsonCConvert(message).(map[string]interface{}), SetQueryOptions(options))
	return goToCBoolResult(res, err)
}

//export kuzzle_wrapper_collection_set_headers
func kuzzle_wrapper_collection_set_headers(c *C.collection, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		r := replace != 0
		cToGoCollection(c).SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}

	return
}

//export kuzzle_wrapper_collection_truncate
func kuzzle_wrapper_collection_truncate(c *C.collection, options *C.query_options) *C.ack_result {
	res, err := cToGoCollection(c).Truncate(SetQueryOptions(options))
	return goToCAckResult(res, err)
}
