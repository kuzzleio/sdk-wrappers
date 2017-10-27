package main

/*
	#cgo CFLAGS: -I../../headers
	
	#include <stdlib.h>
	#include "kuzzle.h"
	#include "sdk_wrappers_internal.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"encoding/json"
)

// Allocates memory
func goToCKuzzleMeta(gMeta *types.Meta) *C.meta {
	if gMeta == nil {
		return nil
	}

	result := (*C.meta)(C.calloc(1, C.sizeof_meta))
	result.author = C.CString(gMeta.Author)
	result.created_at = C.ulonglong(gMeta.CreatedAt)
	result.updated_at = C.ulonglong(gMeta.UpdatedAt)
	result.deleted_at = C.ulonglong(gMeta.DeletedAt)
	result.updater = C.CString(gMeta.Updater)
	result.active = C.bool(gMeta.Active)

	return result
}

// Allocates memory
func goToCShards(gShards *types.Shards) *C.shards {
	if gShards == nil {
		return nil
	}

	result := (*C.shards)(C.calloc(1, C.sizeof_shards))
	result.failed = C.int(gShards.Failed)
	result.successful = C.int(gShards.Successful)
	result.total = C.int(gShards.Total)

	return result
}

// Allocates memory
func goToCDocument(col *C.collection, gDoc *collection.Document) *C.document {
	result := (*C.document)(C.calloc(1, C.sizeof_document))

	result.id = C.CString(gDoc.Id)
	result.index = C.CString(gDoc.Index)
	result.result = C.CString(gDoc.Result)
	result.collection = C.CString(gDoc.Collection)
	result.meta = goToCKuzzleMeta(gDoc.Meta)
	result.shards = goToCShards(gDoc.Shards)
	result._collection = col

	if string(gDoc.Content) != "" {
		buffer := C.CString(string(gDoc.Content))
		result.content = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))
	} else {
		result.content = C.json_object_new_object()
	}

	result.version = C.int(gDoc.Version)
	result.created = C.bool(gDoc.Created)

	return result
}

// Allocates memory
func goToCDocumentResult(col *C.collection, goRes *collection.Document, err error) *C.document_result {
	result := (*C.document_result)(C.calloc(1, C.sizeof_document_result))

	if err != nil {
		Set_document_error(result, err)
		return result
	}

	result.result = goToCDocument(col, goRes)

	return result
}

// Allocates memory
func goToCAckResult(goRes *types.AckResponse, err error) *C.ack_result {
	result := (*C.ack_result)(C.calloc(1, C.sizeof_ack_result))

	if err != nil {
		Set_ack_result_error(result, err)
		return result
	}

	result.acknowledged = C.bool(goRes.Acknowledged)
	result.shards_acknowledged = C.bool(goRes.ShardsAcknowledged)

	return result
}

// Allocates memory
func goToCStringResult(goRes string, err error) *C.string_result {
	result := (*C.string_result)(C.calloc(1, C.sizeof_string_result))

	if err != nil {
		Set_string_result_error(result, err)
		return result
	}

	result.result = C.CString(goRes)

	return result
}

func goToCStringArrayResult(goRes []string, err error) *C.string_array_result {
	result := (*C.string_array_result)(C.calloc(1, C.sizeof_string_array_result))

	if err != nil {
		Set_string_array_result_error(result, err)
		return result
	}

	result.result = (**C.char)(C.calloc(C.size_t(len(goRes)), C.sizeof_char_ptr))
	result.length = C.ulong(len(goRes))

	cArray := (*[1<<30 - 1]*C.char)(unsafe.Pointer(result.result))[:len(goRes):len(goRes)]

	for i, substring := range goRes {
		cArray[i] = C.CString(substring)
	}

	return result
}

// Allocates memory
func goToCIntResult(goRes int, err error) *C.int_result {
	result := (*C.int_result)(C.calloc(1, C.sizeof_int_result))

	if err != nil {
		Set_int_result_error(result, err)
		return result
	}

	result.result = C.longlong(goRes)

	return result
}

// Allocates memory
func goToCDoubleResult(goRes float64, err error) *C.double_result {
	result := (*C.double_result)(C.calloc(1, C.sizeof_double_result))

	if err != nil {
		Set_double_result_error(result, err)
		return result
	}

	result.result = C.double(goRes)

	return result
}

// Allocates memory
func goToCBoolResult(goRes bool, err error) *C.bool_result {
	result := (*C.bool_result)(C.calloc(1, C.sizeof_bool_result))

	if err != nil {
		Set_bool_result_error(result, err)
		return result
	}

	result.result = C.bool(goRes)

	return result
}

// Allocates memory
func goToCSearchResult(col *C.collection, goRes *collection.SearchResult, err error) *C.search_result {
	result := (*C.search_result)(C.calloc(1, C.sizeof_search_result))

	if err != nil {
		Set_search_result_error(result, err)
		return result
	}

	result.result = (*C.document_search)(C.calloc(1, C.sizeof_document_search))
	result.result.length = C.uint(len(goRes.Hits))
	result.result.total = C.uint(goRes.Total)
	result.result.scrollId = C.CString(goRes.ScrollId)

	if len(goRes.Hits) > 0 {
		result.result.hits = (**C.document)(C.calloc(C.size_t(len(goRes.Hits)), C.sizeof_document_ptr))
		cArray := (*[1<<30 - 1]*C.document)(unsafe.Pointer(result.result.hits))[:len(goRes.Hits):len(goRes.Hits)]

		for i, doc := range goRes.Hits {
			cArray[i] = goToCDocument(col, doc)
		}
	}

	return result
}

// Allocates memory
func goToCMapping(c *C.collection, goMapping *collection.Mapping) *C.mapping {
	result := (*C.mapping)(C.calloc(1, C.sizeof_mapping))

	result.collection = c
	r, _ := json.Marshal(goMapping.Mapping)
	buffer := C.CString(string(r))
	result.mapping = C.json_tokener_parse(buffer)
	C.free(unsafe.Pointer(buffer))

	return result
}

// Allocates memory
func goToCMappingResult(c *C.collection, goRes *collection.Mapping, err error) *C.mapping_result {
	result := (*C.mapping_result)(C.calloc(1, C.sizeof_mapping_result))

	if err != nil {
		Set_mapping_result_error(result, err)
		return result
	}

	result.result = goToCMapping(c, goRes)

	return result
}

// Allocates memory
func goToCSpecification(goSpec *types.Specification) *C.specification {
	result := (*C.specification)(C.calloc(1, C.sizeof_specification))

	result.strict = C.bool(goSpec.Strict)

	f, _ := json.Marshal(goSpec.Fields)
	v, _ := json.Marshal(goSpec.Validators)
	bufferFields := C.CString(string(f))
	bufferValidators := C.CString(string(v))

	result.fields = C.json_tokener_parse(bufferFields)
	result.validators = C.json_tokener_parse(bufferValidators)

	C.free(unsafe.Pointer(bufferFields))
	C.free(unsafe.Pointer(bufferValidators))

	return result
}

// Allocates memory
func goToCSpecificationEntry(goEntry *types.SpecificationEntry) *C.specification_entry {
	result := (*C.specification_entry)(C.calloc(1, C.sizeof_specification_entry))
	result.index = C.CString(goEntry.Index)
	result.collection = C.CString(goEntry.Collection)
	result.validation = goToCSpecification(goEntry.Validation)

	return result
}

// Allocates memory
func goToCSpecificationResult(goRes *types.Specification, err error) *C.specification_result {
	result := (*C.specification_result)(C.calloc(1, C.sizeof_specification_result))

	if err != nil {
		Set_specification_result_err(result, err)
		return result
	}

	result.result = goToCSpecification(goRes)

	return result
}

// Allocates memory
func goToCSpecificationSearchResult(goRes *types.SpecificationSearchResult, err error) *C.specification_search_result {
	result := (*C.specification_search_result)(C.calloc(1, C.sizeof_specification_search_result))

	if err != nil {
		Set_specification_search_result_error(result, err)
		return result
	}

	result.result = (*C.specification_search)(C.calloc(1, C.sizeof_specification_search))
	result.result.length = C.uint(len(goRes.Hits))
	result.result.total = C.uint(goRes.Total)
	result.result.scrollId = C.CString(goRes.ScrollId)

	if len(goRes.Hits) > 0 {
		result.result.hits = (**C.specification_entry)(C.calloc(C.size_t(len(goRes.Hits)), C.sizeof_specification_entry_ptr))
		cArray := (*[1<<30 - 1]*C.specification_entry)(unsafe.Pointer(result.result.hits))[:len(goRes.Hits):len(goRes.Hits)]

		for i, spec := range goRes.Hits {
			cArray[i] = goToCSpecificationEntry(&spec.Source)
		}
	}

	return result
}

func goToCJsonResult(goRes interface{}, err error) *C.json_result {
	result := (*C.json_result)(C.calloc(1, C.sizeof_json_result))

  if err != nil {
    Set_json_result_error(result, err)
    return result
  }

	r, _ := json.Marshal(goRes)

  buffer := C.CString(string(r))

  result.result = C.json_tokener_parse(buffer)

	C.free(unsafe.Pointer(buffer))
  return result
}

func goToCJsonArrayResult(goRes []interface{}, err error) *C.json_array_result {
	result := (*C.json_array_result)(C.calloc(1, C.sizeof_json_array_result))

	if err != nil {
	  Set_json_array_result_error(result, err)
	  return result
	}

	result.length = C.uint(len(goRes))
	result.result = (**C.json_object)(C.calloc(C.size_t(result.length), C.sizeof_json_object_ptr))
	cArray := (*[1<<30 - 1]*C.json_object)(unsafe.Pointer(result.result))[:len(goRes):len(goRes)]

	for i, res := range(goRes) {
		r, _ := json.Marshal(res)
		buffer := C.CString(string(r))
		cArray[i] = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))
	}

	return result
}
