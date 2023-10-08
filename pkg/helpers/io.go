package helpers

import (
	"fmt"
	"io"
	"os"
)

func BytesFromFilePath(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	fileData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer file.Close()

	return fileData, err
}
