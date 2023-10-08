package helpers

/*
#include <stdlib.h>
#include <stdint.h>
*/
import "C"
import (
	"unsafe"
)

type CBytes struct {
	Pointer unsafe.Pointer
	Size    int
}

func (CBytes) Create(bytes []byte) CBytes {
	return CBytes{
		Pointer: C.CBytes(bytes),
		Size:    len(bytes),
	}
}
