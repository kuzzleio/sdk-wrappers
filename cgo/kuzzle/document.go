package main

/*
	#cgo CFLAGS: -I../../headers
	#include "kuzzle.h"
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"github.com/kuzzleio/sdk-go/types"
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
func gotCShards(gShards *types.Shards) *C.shards {
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
	result.shards = gotCShards(gDoc.Shards)

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