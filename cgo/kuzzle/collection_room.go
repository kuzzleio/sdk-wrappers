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
)

//export kuzzle_wrapper_collection_new_room
func kuzzle_wrapper_collection_new_room(room *C.room, c *C.collection, options *C.room_options) {
	var opts types.RoomOptions
	if options != nil {
		opts = SetRoomOptions(options)
	}

	instance := collection.NewRoom((*collection.Collection)(c.instance), opts)

	room.instance = unsafe.Pointer(instance)
}

//export kuzzle_wrapper_collection_get_room_id
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