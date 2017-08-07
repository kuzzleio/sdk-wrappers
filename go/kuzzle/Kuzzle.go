package main

/*
	typedef struct {
	} kuzzle;

	extern char* kuzzle_wrapper_connect();
 */
import "C"
import (
	"github.com/kuzzleio/sdk-go/kuzzle"
	"github.com/kuzzleio/sdk-go/connection/websocket"
	"github.com/kuzzleio/sdk-go/connection"
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

func main() {

}
