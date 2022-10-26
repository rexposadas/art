package util

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func PathToImageJPG(path string) image.Image {
	imageFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	result, err := jpeg.Decode(imageFile)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer imageFile.Close()

	return result
}

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

func Merge(bgFile, signatureFile string) {
	bgFile = "output/art.png"
	signatureFile = "input/signature.png"

	image1, err := os.Open(bgFile)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	actualArt, err := png.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	watermark, err := os.Open(signatureFile)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	signature, err := png.Decode(watermark)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer watermark.Close()

	offset := image.Pt(700, 950)
	b := actualArt.Bounds()
	image3 := image.NewRGBA(b)

	draw.Draw(image3, b, actualArt, image.ZP, draw.Src)

	//mask := image.NewUniform(color.Alpha{A: 128})
	//
	//draw.DrawMask(image3, signature.Bounds().Add(offset), signature,
	//	image.Point{X: -100, Y: -100}, nil, image.Point{X: -200, Y: -200}, draw.Over)

	draw.Draw(image3, signature.Bounds().Add(offset), signature, image.ZP, draw.Over)

	f := fmt.Sprintf("output/merged/%s.png", RnID())
	third, err := os.Create(f)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer third.Close()
}
