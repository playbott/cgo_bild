package image

import (
	"bytes"
	"image"
)

func FromBytes(imageData []byte) (image.Image, error) {
	reader := bytes.NewReader(imageData)
	img, _, err := image.Decode(reader)
	if err != nil {
		return img, err
	}
	return img, nil
}
