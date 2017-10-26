package main

/*
	#cgo CFLAGS: -I../../headers
  #include <stdlib.h>
	#include "kuzzle.h"
*/
import "C"
import (
	"encoding/json"
	"github.com/kuzzleio/sdk-go/kuzzle"
	"unsafe"
)

//export kuzzle_wrapper_get_headers
func kuzzle_wrapper_get_headers(k *C.kuzzle) *C.json_object {
	res := (*kuzzle.Kuzzle)(k.instance).GetHeaders()
	r, _ := json.Marshal(res)

  buffer := C.CString(string(r))
  defer C.free(unsafe.Pointer(buffer))
	
  return C.json_tokener_parse(buffer)
}
