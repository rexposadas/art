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
		common.Citrus,
		common.DarkPink,
	}

	return schemes[rand.Intn(len(schemes))]
}

func RnColor() color.RGBA {
	schemes := []color.RGBA{
		common.Black,
		common.Azure,
		common.PaleTurquoise,
		common.Bisque,
		common.MediumAquamarine,
		common.Plum,
		common.Orange,
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
		common.AntiqueWhite,
	}

	return bg[rand.Intn(len(bg))]
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
