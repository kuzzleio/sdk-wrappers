package main

/*
	#cgo CFLAGS: -I../../headers
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
	buffer := C.CString(json.Marshal(goMapping.Mapping))
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

	bufferFields := C.Cstring(json.Marshal(goSpec.Fields))
	bufferValidators := C.Cstring(json.Marshal(goSpec.Validators))

	result.fields = C.json_tokener_parse(bufferFields)
	result.validators = C.json_tokener_parse(bufferValidators)

	C.free(unsafe.Pointer(bufferFields))
	C.free(unsafe.Pointer(bufferValidators))

	return result
}

// Allocates memory
func goToCSpecificationEntry(goEntry *types.SpecificationEntry) *C.specification_entry {
	result := (*C.specification_entry)(C.calloc(1, C.sizeof_specification_entry))
	result.index = C.Cstring(goEntry.Index)
	result.collection = C.CString(goEntry.Collection)
	result.validation = goToCSpecification(goEntry.Validation)

	return result
}

// Allocates memory
func goToCSpecificationResult(goRes *types.Specification, err error) *C.specification_result {
	result := (*C.specification_result)(C.calloc(1, C.sizeof_specification_result))

	if err != nil {
		Set_specification_result_err(result, err)
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

	result.result = (*C.specification_search)(C.calloc(1, C.specification_search))
	result.result.length = C.uint(len(goRes.Hits))
	result.result.total = C.uint(goRes.Total)
	result.result.scrollId = C.CString(goRes.ScrollId)

	if len(goRes.Hits) > 0 {
		result.result.hits = (**C.specification_entry)(C.calloc(C.size_t(len(goRes.Hits)), C.sizeof_specification_entry_ptr))
		cArray := (*[1<<30 - 1]*C.document)(unsafe.Pointer(result.result.hits))[:len(goRes.Hits):len(goRes.Hits)]

		for i, spec := range goRes.Hits {
			cArray[i] = goToCSpecificationEntry(&spec.Source)
		}
	}

	return result
}
