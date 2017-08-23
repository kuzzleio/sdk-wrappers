package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
	#include <string.h>
*/
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/connection"
	"unsafe"
	"github.com/kuzzleio/sdk-go/types"
	"fmt"
)

var KuzzleInstance *kuzzle.Kuzzle

//export Kuzzle
func Kuzzle(host, protocol *C.char) *C.kuzzle {
	var c connection.Connection

	if C.GoString(protocol) == "websocket" {
		c = websocket.NewWebSocket(C.GoString(host), nil)
	} else {
		return nil
	}

	KuzzleInstance, _ = kuzzle.NewKuzzle(c, nil)

	return &C.kuzzle{}
}

//export kuzzle_wrapper_connect
func kuzzle_wrapper_connect() *C.char {
	err := KuzzleInstance.Connect()
	if err != nil {
		return C.CString(err.Error())
	}

	return nil
}

//export kuzzle_wrapper_get_offline_queue
func kuzzle_wrapper_get_offline_queue() *C.offline_queue {
	//todo does not work
	po := make([]types.QueryObject, 2)

	po[0] = types.QueryObject{RequestId: "42"}

	arr := make([]C.query_object, len(po))
	unsafePointer := unsafe.Pointer(C.malloc(C.size_t(len(po))))

	for _, qo := range po {
		struct_qo := C.query_object{}

		struct_qo.query = C.CString(string(qo.Query))
		struct_qo.requestId = *(*[36]C.char)(unsafe.Pointer(C.CString(qo.RequestId)))
		struct_qo.timestamp = *(*[11]C.char)(unsafe.Pointer(C.CString(qo.Timestamp.String())))

		C.memcpy(unsafePointer, unsafe.Pointer(&struct_qo), C.size_t(unsafe.Sizeof(struct_qo)))
		arr = append(arr, struct_qo)
	}

	fmt.Printf("%s\n", (*unsafe.Pointer)(unsafe.Pointer((*C.query_object)(unsafePointer))))

	oq := C.offline_queue{}
	oq.query = (**C.query_object)(unsafePointer)
	return &oq
}

//export kuzzle_wrapper_get_jwt
func kuzzle_wrapper_get_jwt() *C.char {
	return C.CString(KuzzleInstance.GetJwt())
}

func main() {

}
