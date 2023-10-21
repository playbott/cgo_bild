package encode

import "C"
import (
	"bytes"
	"fmt"
	"github.com/anthonynsimon/bild/imgio"
	"image"
	"image/color"
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

func BMP(img image.Image) ([]byte, error) {
	var byteBuffer bytes.Buffer
	encoder := imgio.BMPEncoder()
	err := encoder(&byteBuffer, img)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	imageData := byteBuffer.Bytes()

	return imageData, nil
}

func PNG24(img image.Image) ([]byte, error) {
	var byteBuffer bytes.Buffer

	out := image.NewRGBA(img.Bounds())

	for y := 0; y < out.Bounds().Dy(); y++ {
		for x := 0; x < out.Bounds().Dx(); x++ {
			col := img.At(x, y).(color.YCbCr)
			// Swap green and red
			r, g, b := color.YCbCrToRGB(col.Y, col.Cb, col.Cr)
			out.SetRGBA(x, y, color.RGBA{
				R: r,
				G: g,
				B: b,
				A: 255,
			})
		}
	}

	encoder := imgio.PNGEncoder()
	err := encoder(&byteBuffer, out)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	imageData := byteBuffer.Bytes()

	return imageData, nil
}

func PNG48(img image.Image) ([]byte, error) {
	var byteBuffer bytes.Buffer

	encoder := imgio.PNGEncoder()
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
