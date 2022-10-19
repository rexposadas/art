package util

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"
)

func Merge(bgFile, fgFile string) {
	image1, err := os.Open(bgFile)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}

	first, err := png.Decode(image1)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer image1.Close()

	watermark, err := os.Open(fgFile)
	if err != nil {
		log.Fatalf("failed to open: %s", err)
	}
	wtrMark, err := png.Decode(watermark)
	if err != nil {
		log.Fatalf("failed to decode: %s", err)
	}
	defer watermark.Close()

	offset := image.Pt(200, 100)
	b := first.Bounds()
	image3 := image.NewRGBA(b)

	draw.Draw(image3, b, first, image.ZP, draw.Src)

	mask := image.NewUniform(color.Alpha{A: 128})

	draw.DrawMask(image3, wtrMark.Bounds().Add(offset), wtrMark,
		image.Point{X: -100, Y: -100}, mask, image.Point{X: -200, Y: -200}, draw.Over)

	f := fmt.Sprintf("output/merged/%s.png", RnID())
	third, err := os.Create(f)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer third.Close()
}
