package samples

import (
	"art/models"
	"art/util"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"math/rand"
	"time"
)

func Perls(cfg *models.Config) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnColor())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	//c.Draw(arts.NewContourLine(500))
	c.Draw(arts.NewContourLine(util.Rn(480, 520)))
	c.ToPNG(cfg.OutURL())
}
