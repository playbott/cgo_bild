package encoder

import "C"
import (
	"github.com/anthonynsimon/bild/imgio"
	"image"
)

func JPEG(image image.Image, outPath string, quality int) error {
	if err := imgio.Save(outPath, image, imgio.JPEGEncoder(quality)); err != nil {
		return err
	}

	return nil
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
