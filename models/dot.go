package models

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/config"
	"log"
	"math/rand"
	"time"
)

func Dot(cfg *config.Config) string {
	if cfg.Canvas.Width < 2080 {
		log.Printf("Canvas width is too small. Minimum is 2080. Got %d", cfg.Canvas.Width)
		return ""
	}

	if cfg.Canvas.Height < 2080 {
		log.Printf("Canvas height is too small. Minimum is 2080. Got %d", cfg.Canvas.Height)
		return ""
	}

	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.SetLineWidth(float64(util.Rn(8, 12)))
	c.SetIterations(15000)
	c.SetColorSchema(util.RnColorScheme())
	c.FillBackground()
	c.Draw(arts.NewDotLine(100, 20, 50, false))

	cfg.Out.Prefix = "dot"
	out := cfg.OutURL()
	if err := c.ToPNG(out); err != nil {
		log.Printf("failed to generate dot images: %s", err)
	}

	return out
}
