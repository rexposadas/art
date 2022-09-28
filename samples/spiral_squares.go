package samples

import (
	"art/models"
	"art/util"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
)

func SpiralSquare(cfg *models.Config) {
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(common.MistyRose)
	c.SetLineWidth(float64(util.Rn(5, 10)))
	c.SetLineColor(util.RnColor())
	c.SetColorSchema(util.RnColorScheme())
	c.SetForeground(util.RnColor())
	c.FillBackground()
	c.Draw(arts.NewSpiralSquare(40, 400, 0.05, true))
	c.ToPNG(cfg.OutURL())
}
