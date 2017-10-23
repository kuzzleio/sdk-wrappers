package main

/*
	#cgo CFLAGS: -I../../headers

	static int sizeArray(char** arr) {
		int i = 0;

		if (!arr || !arr[0])
			return 0;
		while (arr[i])
			i++;

		return i;
	}

	static int sizeDocumentArray(document** arr) {
		int i = 0;

		if (!arr || !arr[0])
			return 0;
		while (arr[i])
			i++;

		return i;
	}

	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

// Allocates memory
func goToCKuzzleMeta(gMeta *types.KuzzleMeta) *C.kuzzle_meta {
	result := (*C.kuzzle_meta)(C.calloc(1, C.sizeof_kuzzle_meta))

	if gMeta != nil {
		result.author = C.CString(gMeta.Author)
		result.created_at = C.ulonglong(gMeta.CreatedAt)
		result.updated_at = C.ulonglong(gMeta.UpdatedAt)
		result.deleted_at = C.ulonglong(gMeta.DeletedAt)
		result.updater = C.CString(gMeta.Updater)

		if gMeta.Active {
			result.active = 1
		} else {
			result.active = 0
		}
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
func goToCDocument(gDoc *collection.Document) *C.document {
	result := (*C.document)(C.calloc(1, C.sizeof_document))

	result.id = C.CString(gDoc.Id)
	result.index = C.CString(gDoc.Index)
	result.result = C.CString(gDoc.Result)
	result.collection = C.CString(gDoc.Collection)
	result.meta = goToCKuzzleMeta(gDoc.Meta)
	result.shards = goToCShards(gDoc.Shards)

	if string(gDoc.Content) != "" {
		buffer := C.CString(string(gDoc.Content))
    result.content = C.json_tokener_parse(buffer)
    C.free(buffer)
	} else {
		result.content = C.json_object_new_object()
	}

	result.version = C.int(gDoc.Version)

	if gDoc.Created {
		result.created = 1
	}

	return result
}

// Allocates memory
func goToCSearchResult(goRes *collection.SearchResult, err error) *C.kuzzle_search_result {
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
			cArray[i] = goToCDocument(doc)
		}
	}

	return result
}

// TODO check if it is still legit - refactor
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

// convert a C char** to a go array of string
func cToGoStrings(arr **C.char, len C.uint) []string {
	if len == 0 {
		return nil
	}

	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(arr))[:len:len]
	goStrings := make([]string, len)
	for i, s := range tmpslice {
		goStrings[i] = C.GoString(s)
	}

	return goStrings
}

// Helper to convert a C document** to a go array of document pointers
// TODO Refactor document
func cToGoDocuments(docs **C.document) []*collection.Document {
	length := C.sizeDocumentArray(docs)
	if length == 0 {
		return nil
	}
	tmpslice := (*[1 << 30]*C.document)(unsafe.Pointer(docs))[:length:length]
	godocuments := make([]*collection.Document, length)
	for i, doc := range tmpslice {
		godocuments[i] = cToGoDocument(doc)
	}
	return godocuments
}

// TODO
func cToGoDocument(doc *C.document) *collection.Document {

}