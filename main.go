package main

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"

import (
	"github.com/anthonynsimon/bild/imgio"
)

func main() {
}

//export JPEGEncode
func JPEGEncode(cInPath *C.char, cOutPath *C.char, cQuality C.int) C.int {
	inPath := C.GoString(cInPath)
	outPath := C.GoString(cOutPath)
	quality := CIntToGo(cQuality)

	image, err := imgio.Open(inPath)
	if err != nil {
		return C.int(-1)
	}

	if err := imgio.Save(outPath, image, imgio.JPEGEncoder(quality)); err != nil {
		return C.int(-2)
	}

	return C.int(0)
}

func CIntToGo(n C.int) int {
	var y *C.int = (*C.int)(C.malloc(C.sizeof_int))
	*y = n
	f := int(*y)
	return f
}
