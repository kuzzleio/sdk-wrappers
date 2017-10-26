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
	"github.com/kuzzleio/sdk-go/security"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

func cToGoControllers(c *C.controllers) (*types.Controllers, error) {
	if c == nil {
		return nil, nil
	}

	if JsonCType(c) != C.json_type_object {
		return nil, types.NewError("Invalid controller structure", 400)
	}
	j := JsonCConvert(c)
	controllers := &types.Controllers{}

	for controller, val := range j.(map[string]interface{}) {
		rawController, ok := val.(map[string]interface{})
		if !ok {
			return nil, types.NewError("Invalid controllers structure", 400)
		}

		rawActions, ok := rawController["Actions"]
		if !ok {
			return nil, types.NewError("Invalid controllers structure", 400)
		}

		actionsMap, ok := rawActions.(map[string]interface{})
		if !ok {
			return nil, types.NewError("Invalid controllers structure", 400)
		}

		controllers.Controllers[controller] = &types.Controller{}
		for action, value := range actionsMap {
			boolValue, ok := value.(bool)
			if !ok {
				return nil, types.NewError("Invalid controllers structure", 400)
			}

			controllers.Controllers[controller].Actions[action] = boolValue
		}
	}

	return controllers, nil
}

func cToGoRole(crole *C.role) (*security.Role, error) {
	id := C.GoString(crole.id)
	var controllers *types.Controllers

	if crole.controllers != nil {
		c, err := cToGoControllers(crole.controllers)

		if err != nil {
			return nil, err
		}
		controllers = c
	}

	role := (*kuzzle.Kuzzle)(crole.kuzzle.instance).Security.NewRole(id, controllers)

	return role, nil
}

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
func cToGoDocuments(col *C.collection, docs **C.document, length C.uint) []*collection.Document {
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
	return cToGoCollection(c)
}

// TODO: Test !
func cToGoMapping(cMapping *C.mapping) *collection.Mapping {
	mapping := collection.NewMapping(cToGoCollection(cMapping.collection))
	mapping.Mapping = JsonCConvert(cMapping.mapping).(types.MappingFields)

	return mapping
}

func cToGoDocument(c *C.collection, cDoc *C.document) *collection.Document {
	col := cToGoCollection(c)
	gDoc := col.Document()
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
