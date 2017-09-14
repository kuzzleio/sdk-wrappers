package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
	#include <string.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"unsafe"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"encoding/json"
	"github.com/kuzzleio/sdk-go/types"
)

//export Kuzzle
func Kuzzle(k *C.kuzzle, host, protocol *C.char, options *C.options) {
	var c connection.Connection

	var opts types.Options
	if options != nil {
		opts = SetOptions(options)
	}

	if C.GoString(protocol) == "websocket" {
		c = websocket.NewWebSocket(C.GoString(host), opts)
	}

	instance, _ := kuzzle.NewKuzzle(c, opts)
	k.instance = unsafe.Pointer(instance)
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
func kuzzle_wrapper_get_offline_queue(k *C.kuzzle, result *C.offline_queue) {
	offlineQueue := (*kuzzle.Kuzzle)(k.instance).GetOfflineQueue()
	qooo := types.QueryObject{RequestId: "test"}
	*offlineQueue = append(*offlineQueue, qooo)

	cArray := C.malloc(C.size_t(len(*offlineQueue)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	query_objects := (*[1<<30 - 1]*C.query_object)(cArray)

	idx := 0
	for _, queryObject := range *offlineQueue {
		qo := C.query_object{}
		qo.timestamp = C.ulonglong(queryObject.Timestamp.Unix())
		qo.request_id = *(*[36]C.char)(unsafe.Pointer(C.CString(queryObject.RequestId)))
		mquery, _ := json.Marshal(queryObject.Query)
		qo.query = C.json_tokener_parse(C.CString(string(mquery)))
		query_objects[idx] = &qo
		idx += 1
	}
	query_objects[idx] = nil

	*result = C.offline_queue{(**C.query_object)(cArray)}
}

//export kuzzle_wrapper_get_jwt
func kuzzle_wrapper_get_jwt(k *C.kuzzle) *C.char {
	return C.CString((*kuzzle.Kuzzle)(k.instance).GetJwt())
}

func main() {

}
