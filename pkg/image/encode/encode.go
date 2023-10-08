package encode

import "C"
import (
	"bytes"
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"image"
)

func JPEG(img image.Image, quality int) ([]byte, error) {
	var byteBuffer bytes.Buffer
	encoder := imgio.JPEGEncoder(quality)
	err := encoder(&byteBuffer, img)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	imageData := byteBuffer.Bytes()

	return imageData, nil
}

func JPEGFromDisc(inPath string, outPath string, quality int) error {
	img, err := imgio.Open(inPath)
	if err != nil {
		return err
	}
	if err := imgio.Save(outPath, img, imgio.JPEGEncoder(quality)); err != nil {
		return err
	}
	return nil
}

func BMP(image image.Image, outPath string) error {
	if err := imgio.Save(outPath, image, imgio.BMPEncoder()); err != nil {
		return err
	}
	return nil
}

func PNG(image image.Image, outPath string) error {
	if err := imgio.Save(outPath, image, imgio.PNGEncoder()); err != nil {
		return err
	}
	return nil
}
