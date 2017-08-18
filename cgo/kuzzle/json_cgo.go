package main

/*
	#cgo CFLAGS: -I../../headers
	#cgo LDFLAGS: -ljson-c
	#include <kuzzle.h>
 */
import "C"

type JsonParser struct {
	content map[string]interface{}
}

func (parser *JsonParser) get_json_value(key string, jobj *C.json_object, content map[string]interface{}) {
	jtype := C.json_object_get_type(jobj)

	switch jtype {
	case C.json_type_boolean:
		if int(C.json_object_get_boolean(jobj)) == 1 {
			content[key] = true
		} else {
			content[key] = false
		}
		break
	case C.json_type_string:
		content[key] = C.GoString(C.json_object_get_string(jobj))
		break
	case C.json_type_double:
		content[key] = float64(C.json_object_get_double(jobj))
		break
	case C.json_type_int:
		content[key] = int(C.json_object_get_int(jobj))
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

	for entry, entry_next := C.json_object_get_object(jobj).head, C.json_object_get_object(jobj).head; entry != nil; entry = entry_next {
		entry_next = entry.next

		jtype := C.json_object_get_type((*C.json_object)(entry.v))

		switch jtype {
		case C.json_type_boolean, C.json_type_double, C.json_type_int, C.json_type_string:
			parser.get_json_value(C.GoString((*C.char)(entry.k)), (*C.json_object)(entry.v), content)
			break
		case C.json_type_object:
			C.json_object_object_get_ex(jobj, (*C.char)(entry.k), &jobj)
			content[C.GoString((*C.char)(entry.k))] = make(map[string]interface{})
			parser.parse_cjson(jobj, content[C.GoString((*C.char)(entry.k))].(map[string]interface{}))
			break
		}
	}
}

func (parser JsonParser) GetContent() map[string]interface{} {
	return parser.content
}