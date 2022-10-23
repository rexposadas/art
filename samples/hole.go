package samples

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util"
	"math/rand"
	"time"
)

func Hole(cfg *models.Config) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.SetIterations(1200)
	c.Draw(arts.NewPixelHole(60))

	url := cfg.OutURL()
	c.ToPNG(url)

	return url
}
