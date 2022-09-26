package samples

import (
	"art/models"
	"art/util"
	"fmt"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"

	"image/color"
	"math/rand"
	"time"
)

func Circles(cfg *models.Config) {
	rand.Seed(time.Now().Unix())

	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetColorSchema(cfg.Colors.Scheme)
	c.Draw(arts.NewCircleLoop2(util.Rn(5, 8)))
	c.ToPNG(cfg.OutURL())
}

func CircleGradient(filename string) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0x11, 0x60, 0xC6, 0xFF},
		{0xFD, 0xD9, 0x00, 0xFF},
		{0xF5, 0xB4, 0xF8, 0xFF},
		{0xEF, 0x13, 0x55, 0xFF},
		{0xF4, 0x9F, 0x0A, 0xFF},
	}
	c := generativeart.NewCanva(500, 500)
	c.SetBackground(common.White)
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewColorCircle2(30))
	c.ToPNG(fmt.Sprintf("output/%s.png", filename))
}

func oldBackgroup() color.RGBA {
	return color.RGBA{8, 10, 20, 255}
}
