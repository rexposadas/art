package samples

import (
	"github.com/jdxyw/generativeart/common"
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util"
	"math/rand"
	"time"

	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
)

func Smoke(cfg *models.Config) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(common.Black)
	c.SetLineWidth(1.0)
	c.SetLineColor(util.RnColor())
	c.SetAlpha(30)
	c.SetColorSchema(util.RnColorScheme())
	c.SetIterations(4)
	c.FillBackground()
	c.Draw(arts.NewSilkSmoke(400, 20, 0.2, 2, 10, 30, false))
	c.ToPNG(cfg.OutURL())
}
