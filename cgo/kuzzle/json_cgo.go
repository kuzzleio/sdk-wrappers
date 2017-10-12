package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
 */
import "C"
import (
	"unsafe"
)

type JsonParser struct {
	content map[string]interface{}
}

func (parser *JsonParser) get_json_value(key string, jobj *C.json_object, content map[string]interface{}, isArray bool) {
	jtype := C.json_object_get_type(jobj)

	switch jtype {
	case C.json_type_boolean:
		var v bool
		if int(C.json_object_get_boolean(jobj)) == 1 {
			v = true
		} else {
			v = false
		}

		if isArray {
			arr := make([]interface{}, 0)
			if content[key] != nil {
				for _, v := range content[key].([]interface{}) {
					arr = append(arr, v)
				}
			}
			arr = append(arr, v)
			content[key] = arr
		} else {
			content[key] = v
		}
		break
	case C.json_type_string:
		if isArray {
			arr := make([]interface{}, 0)
			if content[key] != nil {
				for _, v := range content[key].([]interface{}) {
					arr = append(arr, v)
				}
			}
			arr = append(arr, C.GoString(C.json_object_get_string(jobj)))
			content[key] = arr
		} else {
			content[key] = C.GoString(C.json_object_get_string(jobj))
		}
		break
	case C.json_type_double:
		if isArray {
			arr := make([]interface{}, 0)
			if content[key] != nil {
				for _, v := range content[key].([]interface{}) {
					arr = append(arr, v)
				}
			}
			arr = append(arr, float64(C.json_object_get_double(jobj)))
			content[key] = arr
		} else {
			content[key] = float64(C.json_object_get_double(jobj))
		}
		break
	case C.json_type_int:
		if isArray {
			arr := make([]interface{}, 0)
			if content[key] != nil {
				for _, v := range content[key].([]interface{}) {
					arr = append(arr, v)
				}
			}
			arr = append(arr, int(C.json_object_get_int(jobj)))
			content[key] = arr
		} else {
			content[key] = int(C.json_object_get_int(jobj))
		}
		break
	}
}

func (parser *JsonParser) Parse(jobj *C.json_object) {
	if parser.content == nil {
		parser.content = make(map[string]interface{})
	}
	parser.parse_cjson(jobj, parser.content)
}

func (parser *JsonParser) parse_cjson(jobj *C.json_object, content map[string]interface{}) {
	if jobj == nil {
		return
	}

	obj := C.json_object_get_object(jobj)
	if obj == nil {
		return
	}

	for entry, entry_next := obj.head, obj.head; entry != nil; entry = entry_next {
		entry_next = entry.next

		key := (*C.char)(entry.k)
		value := (*C.json_object)(entry.v)

		jtype := C.json_object_get_type(value)

		switch jtype {
		case C.json_type_boolean, C.json_type_double, C.json_type_int, C.json_type_string:
			parser.get_json_value(C.GoString(key), value, content, false)
			break
		case C.json_type_object:
			C.json_object_object_get_ex(jobj, key, &jobj)
			content[C.GoString(key)] = make(map[string]interface{})
			parser.parse_cjson(jobj, content[C.GoString(key)].(map[string]interface{}))
			break
		case C.json_type_array:
			parser.parseArray(jobj, key, content)
			break
		}
	}
}

func (parser *JsonParser) parseArray(jobj *C.json_object, key *C.char, content map[string]interface{}) {
	jarray := jobj

	if key != nil {
		C.json_object_object_get_ex(jobj, key, &jarray)
	}

	arraylen := C.json_object_array_length(jarray)
	var jvalue *C.json_object

	for i := 0; i < int(arraylen); i++ {
		jvalue = C.json_object_array_get_idx(jarray, C.int(i))
		jtype := C.json_object_get_type(jvalue)
		if jtype == C.json_type_array {
			parser.parseArray(jvalue, nil, content)
		} else if jtype != C.json_type_object {
			parser.get_json_value(C.GoString(key), jvalue, content, true)
		} else {
			parser.parse_cjson(jvalue, content)
		}
	}
}

func (parser JsonParser) GetContent() map[string]interface{} {
	return parser.content
}

//export kuzzle_wrapper_json_new
func kuzzle_wrapper_json_new(jobj *C.JsonObject) {
	jobj.jobj = C.json_object_new_object()
}

//export kuzzle_wrapper_json_put
func kuzzle_wrapper_json_put(jobj *C.json_object, key *C.char, content unsafe.Pointer, kind C.int) {
	if kind == 0 {
		//string
		C.json_object_object_add(jobj, key, C.json_object_new_string((*C.char)(content)))
	} else if kind == 1 {
		//int
		C.json_object_object_add(jobj, key, C.json_object_new_int64((C.int64_t)(*(*C.int)(content))))
	} else if kind == 2 {
		//double
		C.json_object_object_add(jobj, key, C.json_object_new_double(*(*C.double)(content)))
	} else if kind == 3 {
		//bool
		C.json_object_object_add(jobj, key, C.json_object_new_boolean((C.json_bool)(*(*C.uchar)(content))))
	} else if kind == 4 {
		//json_object
		C.json_object_object_add(jobj, key, ((*C.JsonObject)(content)).jobj)
	}
}

//export kuzzle_wrapper_json_get_string
func kuzzle_wrapper_json_get_string(jobj *C.json_object, key *C.char) *C.char {
	value := C.json_object_new_object()
	C.json_object_object_get_ex(jobj, key, &value)

	return C.json_object_get_string(value)
}

//export kuzzle_wrapper_json_get_int
func kuzzle_wrapper_json_get_int(jobj *C.json_object, key *C.char) C.int {
	value := C.json_object_new_object()
	C.json_object_object_get_ex(jobj, key, &value)

	return C.int(C.json_object_get_int64(value))
}

//export kuzzle_wrapper_json_get_double
func kuzzle_wrapper_json_get_double(jobj *C.json_object, key *C.char) C.double {
	value := C.json_object_new_object()
	C.json_object_object_get_ex(jobj, key, &value)

	return C.json_object_get_double(value)
}

//export kuzzle_wrapper_json_get_bool
func kuzzle_wrapper_json_get_bool(jobj *C.json_object, key *C.char) C.json_bool {
	value := C.json_object_new_object()
	C.json_object_object_get_ex(jobj, key, &value)

	return C.json_object_get_boolean(value)
}

//export kuzzle_wrapper_json_get_json_object
func kuzzle_wrapper_json_get_json_object(jobj *C.json_object, key *C.char) C.JsonObject {
	value := C.json_object_new_object()
	C.json_object_object_get_ex(jobj, key, &value)

	jo := C.JsonObject{}
	jo.jobj = value
	return jo
}
