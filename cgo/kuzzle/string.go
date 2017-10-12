package main

import "C"

func ToCString_2048(str string) [2048]C.char {
	var cString [2048]C.char

	for i, v := range str {
		cString[i] = (C.char)(v)
	}

	return cString
}

func ToCString_36(str string) [36]C.char {
	var cString [36]C.char

	for i, v := range str {
		cString[i] = (C.char)(v)
	}

	return cString
}

func ToCString_128(str string) [128]C.char {
	var cString [128]C.char

	for i, v := range str {
		cString[i] = (C.char)(v)
	}

	return cString
}

func ToCString_512(str string) [512]C.char {
	var cString [512]C.char

	for i, v := range str {
		cString[i] = (C.char)(v)
	}

	return cString
}