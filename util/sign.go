package util

import (
	"fmt"
	"github.com/fogleman/gg"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"log"
	"os"
)

// Sign the art.
// Provide an image and a signature and returns a new image with the signature.
func Sign(pic, signature image.Image) image.Image {

	offset := image.Pt(700, 950)
	b := pic.Bounds()
	image3 := image.NewRGBA(b)

	draw.Draw(image3, b, pic, image.ZP, draw.Src)
	draw.Draw(image3, signature.Bounds().Add(offset), signature, image.ZP, draw.Over)

	f := fmt.Sprintf("output/merged/%s.png", RnID())
	third, err := os.Create(f)
	if err != nil {
		log.Fatalf("failed to create: %s", err)
	}
	jpeg.Encode(third, image3, &jpeg.Options{Quality: jpeg.DefaultQuality})
	defer third.Close()

	return image3
}

// SignWithText uses text to sign the art.
// See for fond: https://www.1001freefonts.com/golang.font
func SignWithText(org image.Image, color color.Color, url string) {
	dc := gg.NewContextForImage(org)
	dc.SetColor(color)
	if err := dc.LoadFontFace("fonts/always-in-my-heart.ttf", 34); err != nil {
		panic(err)
	}
	dc.DrawStringAnchored("Rex Posadas", 1000, 1000, 0.5, 0.5)
	dc.DrawStringAnchored("2022", 1000, 1030, 0.5, 0.5)

	if err := dc.SavePNG(url); err != nil {
		log.Printf("failed to sign image: %s", err)
	}
}
