package models

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/config"
	"log"

	"math/rand"
	"time"
)

func Circles(cfg *config.Config) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(common.Black)
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.Draw(arts.NewCircleLoop2(util.Rn(5, 8)))

	cfg.Out.Prefix = "circles"
	url := cfg.OutURL()
	if err := c.ToPNG(url); err != nil {
		log.Printf("failed to generate circle gradient images: %s", err)
	}

	return url
}

func CirclesGrid(cfg *config.Config) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnColor())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.SetLineWidth(2.0)
	c.Draw(arts.NewCircleGrid(4, 6))

	cfg.Out.Prefix = "grid"

	url := cfg.OutURL()
	if err := c.ToPNG(url); err != nil {
		log.Printf("failed to generate circle gradient images: %s", err)
	}

	return url
}

func CircleGradient(cfg *config.Config) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnColor())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	c.Draw(arts.NewColorCircle2(util.Rn(25, 35)))

	cfg.Out.Prefix = "circle-gradient"
	out := cfg.OutURL()
	if err := c.ToPNG(out); err != nil {
		log.Printf("failed to generate circle gradient images: %s", err)
		return ""
	}

	return out
}

func CirclesLoop(cfg *config.Config) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.SetLineWidth(1)
	c.SetLineColor(util.RnColor())
	c.SetAlpha(30)
	c.SetIterations(1000)
	c.FillBackground()
	c.Draw(arts.NewCircleLoop(100))
	if err := c.ToPNG(cfg.OutURL()); err != nil {
		log.Printf("failed to generate circle gradient images: %s", err)
	}
}
