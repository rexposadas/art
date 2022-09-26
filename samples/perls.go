package samples

import (
	"art/models"
	"art/util"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"math/rand"
	"time"
)

func Perls(cfg *models.Config) {
	rand.Seed(time.Now().Unix())

	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.SetAlpha(120)
	c.SetLineWidth(0.3)
	c.FillBackground()
	c.SetIterations(200)
	c.SetColorSchema(cfg.Colors.Scheme)
	c.Draw(arts.NewPerlinPerls(10, 200, 40, 80))
	c.ToPNG(cfg.OutURL())
}
