package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

// Allocates memory
func goToCKuzzleMeta(gMeta *types.Meta) *C.meta {
	result := (*C.meta)(C.calloc(1, C.sizeof_meta))

	if gMeta != nil {
		result.author = C.CString(gMeta.Author)
		result.created_at = C.ulonglong(gMeta.CreatedAt)
		result.updated_at = C.ulonglong(gMeta.UpdatedAt)
		result.deleted_at = C.ulonglong(gMeta.DeletedAt)
		result.updater = C.CString(gMeta.Updater)
		result.active = C.bool(gMeta.Active)
	}

	return result
}

// Allocates memory
func goToCShards(gShards *types.Shards) *C.shards {
	result := (*C.shards)(C.calloc(1, C.sizeof_shards))

	if gShards != nil {
		result.failed = C.int(gShards.Failed)
		result.successful = C.int(gShards.Successful)
		result.total = C.int(gShards.Total)
	}

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
func goToCSearchResult(col *C.collection, goRes *collection.SearchResult, err error) *C.kuzzle_search_result {
	result := (*C.kuzzle_search_result)(C.calloc(1, C.sizeof_kuzzle_search_result))

	if err != nil {
		Set_kuzzle_search_result_error(result, err)
		return result
	}

	result.result = (*C.search_result)(C.calloc(1, C.sizeof_search_result))
	result.result.total = C.int(goRes.Total)
	result.result.scrollId = C.CString(goRes.ScrollId)

	if len(goRes.Hits) > 0 {
		result.result.hits = (**C.document)(C.calloc(C.size_t(goRes.Total), C.sizeof_char_ptr))
		cArray := (*[1<<30 - 1]*C.document)(unsafe.Pointer(result.result.hits))[:goRes.Total:goRes.Total]

		for i, doc := range goRes.Hits {
			cArray[i] = goToCDocument(col, doc)
		}
	}

	return result
}

/*
  TODO: Must be re-done
func goToCSpecificationSearchResult(goRes *types.KuzzleSpecificationSearchResult, cRes *C.kuzzle_specification_search_result) {
	cRes.result.total = C.int(goRes.Total)

	if len(goRes.Hits) > 0 {
		hits := make([]*C.kuzzle_specification, len(goRes.Hits) + 1)

		for i := 0; i < len(goRes.Hits); i++ {
			var spec C.kuzzle_specification
			// TODO register it in global
			t := goRes.Hits[i]
			spec.instance = unsafe.Pointer(&t)
			hits[i] = &spec
		}
		hits[len(goRes.Hits)] = nil

		cRes.result.hits = &hits[0]
	}
}

 */