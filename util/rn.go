package util

import (
	"github.com/google/uuid"
	"image/color"
	"math/rand"
	"strings"
	"time"
)

func RandomBackGround() color.RGBA {
	return color.RGBA{
		R: Rn8(4, 10),
		G: Rn8(4, 10),
		B: Rn8(4, 10),
		A: Rn8(100, 255),
	}
}

func Rn8(min, max int) uint8 {
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(max-min+1) + min)
}

func Rn(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RnID() string {
	s := uuid.NewString()
	parts := strings.Split(s, "-")
	return parts[0]
}
