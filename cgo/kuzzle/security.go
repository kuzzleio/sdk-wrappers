package main

/*
	#cgo CFLAGS: -I../../headers
	#include <kuzzle.h>
	#include <string.h>
*/
import "C"
import (
  "github.com/kuzzleio/sdk-go/kuzzle"
  "github.com/kuzzleio/sdk-go/security"
  "unsafe"
)

//export Security
func Security(s *C.security, k *C.kuzzle) *C.security {
  instance := security.NewSecurity((*kuzzle.Kuzzle)(k.instance))

  s.instance = unsafe.Pointer(instance)

  return s
}