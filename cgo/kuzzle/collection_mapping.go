package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
)

//export kuzzle_wrapper_new_mapping
func kuzzle_wrapper_new_mapping(c *C.collection) *C.mapping {
	cm := (*C.mapping)(C.calloc(1, C.sizeof_mapping))
	cm.mapping = C.json_object_new_object()
	cm.collection = c

	return cm
}

//export kuzzle_wrapper_mapping_apply
func kuzzle_wrapper_mapping_apply(cm *C.mapping, options *C.query_options) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))
	opts := SetQueryOptions(options)
	_, err := cToGoMapping(cm).Apply(opts)

	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	result.result = C.bool(true)

	return result
}

//export kuzzle_wrapper_mapping_refresh
func kuzzle_wrapper_mapping_refresh(cm *C.mapping, options *C.query_options) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))
	opts := SetQueryOptions(options)
	_, err := cToGoMapping(cm).Refresh(opts)

	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	result.result = C.bool(true)

	return result
}

//export kuzzle_wrapper_mapping_set
func kuzzle_wrapper_mapping_set(cm *C.mapping, jMap *C.json_object) {
	var mappings types.MappingFields

	if JsonCType(jMap) == C.json_type_object {
		jsonString := []byte(C.GoString(C.json_object_to_json_string(jMap)))
		json.Unmarshal(jsonString, &mappings)
	}

	cToGoMapping(cm).Set(&mappings)

	return
}

//export kuzzle_wrapper_mapping_set_headers
func kuzzle_wrapper_mapping_set_headers(cm *C.mapping, content *C.json_object, replace C.uint) {
	if JsonCType(content) == C.json_type_object {
		r := replace != 0
		cToGoMapping(cm).SetHeaders(JsonCConvert(content).(map[string]interface{}), r)
	}

	return
}

