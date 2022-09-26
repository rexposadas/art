package samples

import (
	"art/models"
	"art/util"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"math/rand"
	"time"
)

func Hole(cfg *models.Config) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.SetIterations(1200)
	c.Draw(arts.NewPixelHole(60))
	c.ToPNG(cfg.OutURL())
}
