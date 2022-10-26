package models

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/config"
	"math/rand"
	"time"
)

func julia1(z complex128) complex128 {
	c := complex(-0.1, 0.651)

	z = z*z + c

	return z
}

func Julia(cfg config.Config) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetIterations(800)
	c.SetColorSchema(util.RnColorScheme())
	c.FillBackground()
	c.Draw(arts.NewJulia(julia1, 40, 1.5, 1.5))
	c.ToPNG(cfg.OutURL())
}
