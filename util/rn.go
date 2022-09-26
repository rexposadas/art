package util

import (
	"github.com/google/uuid"
	"github.com/jdxyw/generativeart/common"
	"image/color"
	"math/rand"
	"strings"
	"time"
)

func RnColorScheme() []color.RGBA {
	schemes := [][]color.RGBA{
		common.Youthful,
		common.Cheerful,
		common.DarkRed,
		common.Outdoors,
		common.Sleek,
	}

	return schemes[rand.Intn(len(schemes))]
}

func RnBackground() color.RGBA {
	bg := []color.RGBA{
		common.Black,
		common.White,
		common.Azure,
		common.Aquamarine,
		common.PaleTurquoise,
		common.Plum,
	}

	return bg[rand.Intn(len(bg))]
}

// Rn8 returns a random number for a given range.
func Rn8(min, max int) uint8 {
	rand.Seed(time.Now().UnixNano())
	return uint8(rand.Intn(max-min+1) + min)
}

// Rn returns a random number for a given range.
func Rn(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func RnID() string {
	s := uuid.NewString()
	parts := strings.Split(s, "-")
	return parts[0]
}
