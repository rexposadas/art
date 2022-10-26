package models

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/config"
	"math/rand"
	"time"
)

func Perls(cfg *config.Config) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnColor())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.Draw(arts.NewContourLine(util.Rn(480, 520)))

	cfg.Out.Prefix = "perls"
	url := cfg.OutURL()
	c.ToPNG(url)

	return url
}
