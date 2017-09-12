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
	"github.com/kuzzleio/sdk-go/types"
	"fmt"
	"github.com/kuzzleio/sdk-go/connection"
	"github.com/kuzzleio/sdk-go/connection/websocket"
)

//export Kuzzle
func Kuzzle(k *C.kuzzle, host, protocol *C.char) {
	var c connection.Connection

	if C.GoString(protocol) == "websocket" {
		c = websocket.NewWebSocket(C.GoString(host), nil)
	}

	instance, _ := kuzzle.NewKuzzle(c, nil)
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
func kuzzle_wrapper_get_jwt(k *C.kuzzle) *C.char {
	return C.CString((*kuzzle.Kuzzle)(k.instance).GetJwt())
}

func main() {

}
