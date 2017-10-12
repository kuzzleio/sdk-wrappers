package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>

	static void call(void* f, json_object* res) {
		((void(*)(json_object*))f)(res);
	}
*/
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"unsafe"
)

//export kuzzle_wrapper_add_listener
func kuzzle_wrapper_add_listener(k *C.Kuzzle, e C.int, cb unsafe.Pointer) {
	c := make(chan interface{})

	kuzzle.AddListener((*kuzzle.Kuzzle)(k.instance), int(e), c)
	go func() {
		res := <-c

		var jsonRes *C.json_object
		r, _ := json.Marshal(res)

		jsonRes = C.json_tokener_parse(C.CString(string(r)))

		C.call(cb, jsonRes)
	}()
}
