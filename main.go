package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"cgo_bild/api"
	"unsafe"
)

func main() {}

//export JPEGEncode
func JPEGEncode(b unsafe.Pointer, s C.int, pQuality C.int) (unsafe.Pointer, C.int, C.int) {
	inputFile := C.GoBytes(b, s)
	quality := cGoInt(pQuality)
	ptr, size, errorCode := api.JPEGEncode(inputFile, quality)
	return ptr, C.int(size), C.int(errorCode)
}

func cGoInt(n C.int) int {
	var y *C.int = (*C.int)(C.malloc(C.sizeof_int))
	*y = n
	return int(*y)
}

func cCharFromBytes(b []byte) *C.char {
	// Allocate a contiguous block of memory for the C.char array.
	cCharPtr := C.malloc(C.size_t(len(b) + 1))

	// Return a pointer to the C.char array.
	return (*C.char)(cCharPtr)
}
