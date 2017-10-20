package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include <string.h>
	#include "kuzzle.h"

	typedef query_object *query_object_ptr;
*/
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/types"
	"unsafe"
)

// unregister an instance from the instances map
//export unregisterKuzzle
func unregisterKuzzle(k *C.Kuzzle) {
	delete(instances, (*kuzzle.Kuzzle)(k.instance))
}

//export kuzzle_wrapper_new_kuzzle
func kuzzle_wrapper_new_kuzzle(k *C.Kuzzle, host, protocol *C.char, options *C.Options) {
	var c connection.Connection

	if instances == nil {
		instances = make(map[interface{}]interface{})
	}

	var opts types.Options
	if options != nil {
		opts = SetOptions(options)
	}

	if C.GoString(protocol) == "websocket" {
		c = websocket.NewWebSocket(C.GoString(host), opts)
	}

	inst, _ := kuzzle.NewKuzzle(c, opts)
	registerInstance(inst)

	k.instance = unsafe.Pointer(inst)
}

//export kuzzle_wrapper_connect
func kuzzle_wrapper_connect(k *C.Kuzzle) *C.char {
	err := (*kuzzle.Kuzzle)(k.instance).Connect()
	if err != nil {
		// TODO Must be freed in C
		return C.CString(err.Error())
	}

	return nil
}

//export kuzzle_wrapper_get_offline_queue
func kuzzle_wrapper_get_offline_queue(k *C.Kuzzle) *C.offline_queue {
	result := (*C.offline_queue)(C.calloc(1, C.sizeof_offline_queue))

	offlineQueue := *(*kuzzle.Kuzzle)(k.instance).GetOfflineQueue()
	result.length = C.ulong(len(offlineQueue))

	result.queries = (**C.query_object)(C.calloc(C.size_t(len(offlineQueue)), C.sizeof_query_object_ptr))
	queryObjects := (*[1<<30 - 1]*C.query_object)(unsafe.Pointer(result.queries))[:result.length:result.length]

	idx := 0
	for _, queryObject := range offlineQueue {
		queryObjects[idx] = (*C.query_object)(C.calloc(1, C.sizeof_query_object))
		queryObjects[idx].timestamp = C.ulonglong(queryObject.Timestamp.Unix())
		queryObjects[idx].request_id = C.CString(queryObject.RequestId)
		mquery, _ := json.Marshal(queryObject.Query)

		buffer := C.CString(string(mquery))
		queryObjects[idx].query = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))

		idx += 1
	}

	return result
}

//export kuzzle_wrapper_get_jwt
func kuzzle_wrapper_get_jwt(k *C.Kuzzle) *C.char {
	// TODO Must be freed in C
	return C.CString((*kuzzle.Kuzzle)(k.instance).GetJwt())
}

func main() {

}
