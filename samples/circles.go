package samples

import (
	"art/models"
	"art/util"
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

func CirclesDot(cfg *models.Config) {
	rand.Seed(time.Now().Unix())
	colors := []color.RGBA{
		{0xFF, 0xBE, 0x0B, 0xFF},
		{0xFB, 0x56, 0x07, 0xFF},
		{0xFF, 0x00, 0x6E, 0xFF},
		{0x83, 0x38, 0xEC, 0xFF},
		{0x3A, 0x86, 0xFF, 0xFF},
	}
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnColor())
	c.FillBackground()
	c.SetColorSchema(colors)
	c.Draw(arts.NewDotsWave(util.Rn(250, 350)))
	c.ToPNG(cfg.OutURL())
}

func CirclesGrid(cfg *models.Config) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnColor())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.SetLineWidth(2.0)
	c.Draw(arts.NewCircleGrid(4, 6))
	c.ToPNG(cfg.OutURL())
}

func CircleGradient(cfg *models.Config) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.Draw(arts.NewColorCircle2(util.Rn(25, 35)))
	c.ToPNG(cfg.OutURL())
}
