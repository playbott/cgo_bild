package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"cgo_bild/pkg/encoder"
)

func main() {
}

//export JPEGEncode
func JPEGEncode(cInPath *C.char, cOutPath *C.char, cQuality C.int) C.int {
	inPath := C.GoString(cInPath)
	outPath := C.GoString(cOutPath)
	quality := cGoInt(cQuality)

	err := encoder.JPEGFromDisc(inPath, outPath, quality)
	if err != nil {
		return C.int(-1)
	}
	return C.int(0)
}

func cGoInt(n C.int) int {
	var y *C.int = (*C.int)(C.malloc(C.sizeof_int))
	*y = n
	return int(*y)
}
