package image

import (
	"bytes"
	"fmt"
	"image"
)

func FromBytes(imageData []byte) (image.Image, error) {
	reader := bytes.NewReader(imageData)
	img, _, err := image.Decode(reader)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return img, nil
}
