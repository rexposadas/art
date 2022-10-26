package models

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/config"
	"math/rand"
	"time"
)

func Hole(cfg *config.Config) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.SetIterations(1200)
	c.Draw(arts.NewPixelHole(60))

	cfg.Out.Prefix = "hole"
	url := cfg.OutURL()
	c.ToPNG(url)

	return url
}
