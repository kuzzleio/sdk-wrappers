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
