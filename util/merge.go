package util

import (
	"image"
	"image/png"
	"log"
	"os"
)

func PathToImage(path string) image.Image {
	imageFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	result, err := png.Decode(imageFile)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer imageFile.Close()

	return result
}
