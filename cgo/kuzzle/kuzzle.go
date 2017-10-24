package main

/*
	#cgo CFLAGS: -I../../headers
	#include <stdlib.h>
	#include <string.h>
	#include "kuzzle.h"

	typedef query_object *query_object_ptr;

	static void call(void* f, json_object* res) {
		((void(*)(json_object*))f)(res);
	}
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"unsafe"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"encoding/json"
)

// map which stores instances to keep references in case the gc passes
var instances map[interface{}]interface{}

// register new instance to the instances map
func registerKuzzle(instance interface{}) {
	instances[instance] = nil
}

// unregister an instance from the instances map
//export unregisterKuzzle
func unregisterKuzzle(k *C.kuzzle) {
	delete(instances, (*kuzzle.Kuzzle)(k.instance))
}

//export kuzzle_wrapper_new_kuzzle
func kuzzle_wrapper_new_kuzzle(k *C.kuzzle, host, protocol *C.char, options *C.options) {
	var c connection.Connection

	if instances == nil {
		instances = make(map[interface{}]interface{})
	}

	opts := SetOptions(options)
	
	if C.GoString(protocol) == "websocket" {
		c = websocket.NewWebSocket(C.GoString(host), opts)
	}

	inst, _ := kuzzle.NewKuzzle(c, opts)
	registerKuzzle(inst)

	k.instance = unsafe.Pointer(inst)
}

//export kuzzle_wrapper_connect
func kuzzle_wrapper_connect(k *C.kuzzle) *C.char {
	err := (*kuzzle.Kuzzle)(k.instance).Connect()
	if err != nil {
		return C.CString(err.Error())
	}

	return nil
}

//export kuzzle_wrapper_get_offline_queue
func kuzzle_wrapper_get_offline_queue(k *C.kuzzle) *C.offline_queue {
	result := (*C.offline_queue)(C.calloc(1, C.sizeof_offline_queue))

	offlineQueue := *(*kuzzle.Kuzzle)(k.instance).GetOfflineQueue()
	result.length = C.ulong(len(offlineQueue))

	result.queries = (**C.query_object)(C.calloc(C.size_t(len(offlineQueue)), C.sizeof_query_object_ptr))
	query_objects := (*[1<<30 - 1]*C.query_object)(unsafe.Pointer(result.queries))[:result.length:result.length]

	idx := 0
	for _, queryObject := range offlineQueue {
		query_objects[idx] = (*C.query_object)(C.calloc(1, C.sizeof_query_object))
		query_objects[idx].timestamp = C.ulonglong(queryObject.Timestamp.Unix())
		query_objects[idx].request_id = C.CString(queryObject.RequestId)
		mquery, _ := json.Marshal(queryObject.Query)

		buffer := C.CString(string(mquery))
		query_objects[idx].query = C.json_tokener_parse(buffer)
		C.free(unsafe.Pointer(buffer))

		idx += 1
	}

	return result
}

//export kuzzle_wrapper_disconnect
func kuzzle_wrapper_disconnect(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).Disconnect()
}

//export kuzzle_wrapper_flush_queue
func kuzzle_wrapper_flush_queue(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).FlushQueue()
}

//export kuzzle_wrapper_replay_queue
func kuzzle_wrapper_replay_queue(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).ReplayQueue()
}

//export kuzzle_wrapper_start_queuing
func kuzzle_wrapper_start_queuing(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).StartQueuing()
}

//export kuzzle_wrapper_stop_queuing
func kuzzle_wrapper_stop_queuing(k *C.kuzzle) {
	(*kuzzle.Kuzzle)(k.instance).StopQueuing()
}

//export kuzzle_wrapper_add_listener
func kuzzle_wrapper_add_listener(k *C.kuzzle, e C.int, cb unsafe.Pointer) {
	c := make(chan interface{})

	kuzzle.AddListener((*kuzzle.Kuzzle)(k.instance), int(e), c)
	go func() {
		res := <-c

		var jsonRes *C.json_object
		r, _ := json.Marshal(res)

		buffer := C.CString(string(r))
		defer C.free(unsafe.Pointer(buffer))

		jsonRes = C.json_tokener_parse(buffer)

		C.call(cb, jsonRes)
	}()
}

//export kuzzle_wrapper_remove_listener
func kuzzle_wrapper_remove_listener(k *C.kuzzle, event C.int) {
	(*kuzzle.Kuzzle)(k.instance).RemoveListener(int(event))
}

//export kuzzle_wrapper_get_headers
func kuzzle_wrapper_get_headers(k *C.kuzzle) *C.json_object {
	res := (*kuzzle.Kuzzle)(k.instance).GetHeaders()
	r, _ := json.Marshal(res)

  buffer := C.CString(string(r))
  defer C.free(unsafe.Pointer(buffer))
	
  return C.json_tokener_parse(buffer)
}

//export kuzzle_wrapper_set_headers
func kuzzle_wrapper_set_headers(k *C.kuzzle, content *C.json_object, replace C.uint) {
	jp := JsonParser{}
	jp.Parse(content)

	r := replace != 0
	(*kuzzle.Kuzzle)(k.instance).SetHeaders(jp.GetContent(), r)
}

//export kuzzle_wrapper_set_default_index
func kuzzle_wrapper_set_default_index(k *C.kuzzle, index *C.char) C.int {
	err := (*kuzzle.Kuzzle)(k.instance).SetDefaultIndex(C.GoString(index))
	if err != nil {
		return C.int(C.EINVAL)
	}

	return 0
}

func main() {

}
