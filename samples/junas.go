package samples

import (
	"art/models"
	"art/util"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"image/color"
	"math/rand"
	"time"
)

func Junas(cfg *models.Config) {

	DarkRed := []color.RGBA{
		util.RnColor(),
		util.RnColor(),
		util.RnColor(),
		util.RnColor(),
		util.RnColor(),
		util.RnColor(),
		util.RnColor(),
		util.RnColor(),
		util.RnColor(),
		util.RnColor(),
	}

	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.FillBackground()
	c.SetColorSchema(DarkRed)
	c.SetForeground(util.RnColor())
	c.Draw(arts.NewJanus(10, 0.2))
	c.ToPNG(cfg.OutURL())
}