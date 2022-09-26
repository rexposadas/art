package samples

import (
	"art/models"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"github.com/jdxyw/generativeart/common"
)

func SpiralSquare(cfg *models.Config) {
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(common.MistyRose)
	c.SetLineWidth(10)
	c.SetLineColor(common.Orange)
	c.SetColorSchema(common.DarkRed)
	c.SetForeground(common.Tomato)
	c.FillBackground()
	c.Draw(arts.NewSpiralSquare(40, 400, 0.05, true))
	c.ToPNG(cfg.OutURL())
}
