package samples

import (
	"art/models"
	"art/util"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"math/rand"
	"time"
)

func Dot(cfg *models.Config) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(2080, 2080)
	c.SetBackground(util.RnBackground())
	c.SetLineWidth(float64(util.Rn(8, 12)))
	c.SetIterations(15000)
	c.SetColorSchema(util.RnColorScheme())
	c.FillBackground()
	c.Draw(arts.NewDotLine(100, 20, 50, false))
	c.ToPNG(cfg.OutURL())
}
