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

func BMPEncode(inputFile []byte) (unsafe.Pointer, int, error) {
	img, err := image.FromBytes(inputFile)
	if err != nil {
		return unsafe.Pointer(&inputFile), 1, err
	}
	imageData, err := encode.BMP(img)
	if err != nil {
		fmt.Println(err.Error())
		return unsafe.Pointer(&inputFile), 1, err
	}
	bytes := helpers.CBytes{}.Create(imageData)
	return bytes.Pointer, bytes.Size, nil
}

func PNGEncode(inputFile []byte, colorDepth int) (unsafe.Pointer, int, error) {
	var imageData []byte
	img, err := image.FromBytes(inputFile)
	if err != nil {
		return unsafe.Pointer(&inputFile), 1, err
	}

	switch colorDepth {
	case 24:
		imageData, err = encode.PNG24(img)
	case 48:
		imageData, err = encode.PNG48(img)
	default:
		imageData, err = encode.PNG48(img)
	}

	if err != nil {
		fmt.Println(err.Error())
		return unsafe.Pointer(&inputFile), 1, err
	}

	bytes := helpers.CBytes{}.Create(imageData)
	return bytes.Pointer, bytes.Size, nil
}
