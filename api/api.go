package api

import "C"
import (
	"cgo_bild/pkg/helpers"
	"cgo_bild/pkg/image"
	"cgo_bild/pkg/image/encode"
	"fmt"
	"unsafe"
)

func JPEGEncode(inputFile []byte, quality int) (unsafe.Pointer, int, error) {
	img, err := image.FromBytes(inputFile)
	if err != nil {
		return unsafe.Pointer(&inputFile), 1, err
	}
	imageData, err := encode.JPEG(img, quality)
	if err != nil {
		fmt.Println(err.Error())
		return unsafe.Pointer(&inputFile), 1, err
	}
	bytes := helpers.CBytes{}.Create(imageData)
	return bytes.Pointer, bytes.Size, nil
}
