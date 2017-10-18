package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/collection"
	"unsafe"
	"github.com/kuzzleio/sdk-go/types"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
)

//export kuzzle_wrapper_collection_new_room
// Todo Refactor
func kuzzle_wrapper_collection_new_room(room *C.room, c *C.collection, options *C.room_options) {
	var opts types.RoomOptions
	if options != nil {
		opts = SetRoomOptions(options)
	}

	col := collection.NewCollection((*kuzzle.Kuzzle)(c.kuzzle), C.GoString(c.collection), C.GoString(c.index))
	instance := collection.NewRoom(col, opts)

	room.instance = unsafe.Pointer(instance)
}

//export kuzzle_wrapper_collection_get_room_id
// TODO refactor
func kuzzle_wrapper_collection_get_room_id(room *C.room, result *C.string_result) {
	res := (*collection.Room)(room.instance).GetRoomId()
	result.result = ToCString_2048(res)
}

//export kuzzle_wrapper_collection_get_room_filters
func kuzzle_wrapper_collection_get_room_filters(room *C.room, result *C.json_result) {
	res := (*collection.Room)(room.instance).GetFilters()
	jsonString, err := json.Marshal(res)

	if err == nil {
		cString := C.CString(string(jsonString))
		defer C.free(unsafe.Pointer(cString))
		result.result = C.json_tokener_parse(cString)
	}
}