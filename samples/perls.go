package samples

import (
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util"
	"math/rand"
	"time"
)

func Perls(cfg *models.Config) string {
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
