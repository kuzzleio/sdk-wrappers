package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

//export kuzzle_wrapper_new_collection_mapping
func kuzzle_wrapper_new_collection_mapping(cm *C.collection_mapping, c *C.collection) {
	instance := collection.NewCollectionMapping((*collection.Collection)(c.instance))

	cm.instance = unsafe.Pointer(instance)
}

//export kuzzle_wrapper_collection_mapping_apply
func kuzzle_wrapper_collection_mapping_apply(cm *C.collection_mapping, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	_, err := (*collection.CollectionMapping)(cm.instance).Apply(opts)
	if err != nil {
		cm.error = ToCString_2048(err.Error())
		return 0
	}

	return 0
}

//export kuzzle_wrapper_collection_mapping_refresh
func kuzzle_wrapper_collection_mapping_refresh(cm *C.collection_mapping, options *C.query_options) C.int {
	var opts types.QueryOptions
	if options != nil {
		opts = SetQueryOptions(options)
	}

	_, err := (*collection.CollectionMapping)(cm.instance).Refresh(opts)
	if err != nil {
		cm.error = ToCString_2048(err.Error())
		return 0
	}

	return 0
}

//export kuzzle_wrapper_collection_mapping_set
func kuzzle_wrapper_collection_mapping_set(cm *C.collection_mapping, jMap *C.json_object) {
	mappings := make(types.KuzzleFieldsMapping)

	if JsonCType(jMap) == C.json_type_object {
		for field, mapping := range JsonCConvert(jMap).(map[string]interface{}) {
			f := &types.KuzzleFieldMapping{}
			if mapping.(map[string]interface{})["type"] != nil {
				f.Type = mapping.(map[string]interface{})["type"].(string)
			}
			if mapping.(map[string]interface{})["properties"] != nil {
				f.Properties = mapping.(map[string]interface{})["properties"].(map[string]interface{})
			}
			mappings[field] = f
		}

		(*collection.CollectionMapping)(cm.instance).Set(&mappings)
	}

	return
}

//export kuzzle_wrapper_collection_mapping_set_headers
func kuzzle_wrapper_collection_mapping_set_headers(cm *C.collection_mapping, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		var r bool
		if replace == 1 {
			r = true
		}

		(*collection.CollectionMapping)(cm.instance).SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}

	return
}
