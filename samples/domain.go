package samples

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/config"
)

func cmap(r, m1, m2 float64) color.RGBA {
	rgb := color.RGBA{
		R: uint8(common.Constrain(m1*200*r, 0, 255)),
		G: uint8(common.Constrain(r*200, 0, 255)),
		B: uint8(common.Constrain(m2*255*r, 70, 255)),
		A: 255,
	}
	return rgb
}

func Domain(cfg *config.Config) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnBackground())
	c.FillBackground()
	c.Draw(arts.NewDomainWrap(0.01, 4, 4, 20, cmap))

	cfg.Out.Prefix = "domain"
	url := cfg.OutURL()
	c.ToPNG(url)

	return url
}
