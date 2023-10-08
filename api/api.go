package api

import "C"
import (
	"cgo_bild/pkg/helpers"
	"cgo_bild/pkg/image"
	"cgo_bild/pkg/image/encode"
	"unsafe"
)

func JPEGEncode(inputFile []byte, quality int) (unsafe.Pointer, int, int) {
	img, err := image.FromBytes(inputFile)
	imageData, err := encode.JPEG(img, quality)
	if err != nil {
		return nil, -1, -1
	}
	bytes := helpers.CBytes{}.Create(imageData)
	return bytes.Pointer, bytes.Size, 0
}
