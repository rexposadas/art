package util

import (
	"fmt"
	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
)

func Bild() {

	img, err := imgio.Open("testdata/alexa.png")
	if err != nil {
		fmt.Println(err)
		return
	}

	result := effect.Dilate(img, 3)
	saveImage("testdata/result.png", result)

	result = effect.EdgeDetection(img, 1.0)
	saveImage("testdata/result-edge.png", result)

	result = effect.Emboss(img)
	saveImage("testdata/result-emboss.png", result)

	result = effect.Invert(img)
	saveImage("testdata/result-invert.png", result)

	result = effect.Sobel(img)
	saveImage("testdata/result-sobel.png", result)

}
