package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include "kuzzle.h"
 */
import "C"
import (
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
	"github.com/kuzzleio/sdk-go/collection"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

func cToGoSearchFilters(searchFilters *C.search_filters) *types.SearchFilters {
	return &types.SearchFilters{
		Query: JsonCConvert(searchFilters.query),
		Sort: JsonCConvert(searchFilters.sort).([]interface{}),
		Aggregations: JsonCConvert(searchFilters.aggregations),
		SearchAfter: JsonCConvert(searchFilters.search_after).([]interface{}),
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
func cToGoDocuments( col *C.collection, docs **C.document, length C.uint) []*collection.Document {
	if length == 0 {
		return nil
	}
	tmpslice := (*[1 << 30]*C.document)(unsafe.Pointer(docs))[:length:length]
	godocuments := make([]*collection.Document, length)
	for i, doc := range tmpslice {
		godocuments[i] = cToGoDocument(col, doc)
	}
	return godocuments
}

func cToGoShards(cShards *C.shards) *types.Shards {
	return &types.Shards{
		Total: int(cShards.total),
		Successful: int(cShards.successful),
		Failed: int(cShards.failed),
	}
}

func cToGoKuzzleMeta(cMeta *C.meta) *types.Meta {
	return &types.Meta{
		Author: C.GoString(cMeta.author),
		CreatedAt: int(cMeta.created_at),
		UpdatedAt: int(cMeta.updated_at),
		Updater: C.GoString(cMeta.updater),
		Active: bool(cMeta.active),
		DeletedAt: int(cMeta.deleted_at),
	}
}

func cToGoCollection(c *C.collection) *collection.Collection {
	return collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle.instance), C.GoString(c.collection), C.GoString(c.index))
}

func cToGoMapping(cMapping *C.mapping) *collection.Mapping {
	mapping := collection.NewMapping(cToGoCollection(cMapping.collection))
	json.Unmarshal([]byte(C.GoString(C.json_object_to_json_string(cMapping.mapping))), &mapping.Mapping)

	return mapping
}

func cToGoSpecification(cSpec *C.specification) *types.Specification {
	spec := types.Specification{}
	spec.Strict = bool(cSpec.strict)
	json.Unmarshal([]byte(C.GoString(C.json_object_to_json_string(cSpec.fields))), &spec.Fields)
	json.Unmarshal([]byte(C.GoString(C.json_object_to_json_string(cSpec.validators))), &spec.Validators)

	return &spec
}

func cToGoDocument(c *C.collection, cDoc *C.document) *collection.Document {
	gDoc := cToGoCollection(c).Document()
	gDoc.Id = C.GoString(cDoc.id)
	gDoc.Index = C.GoString(cDoc.index)
	gDoc.Meta = cToGoKuzzleMeta(cDoc.meta)
	gDoc.Shards = cToGoShards(cDoc.shards)
	gDoc.Content = []byte(C.GoString(C.json_object_to_json_string(cDoc.content)))
	gDoc.Version = int(cDoc.version)
	gDoc.Result = C.GoString(cDoc.result)
	gDoc.Collection = C.GoString(cDoc.collection)
	gDoc.Created = bool(cDoc.created)

	return gDoc
}
