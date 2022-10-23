package samples

import (
	"github.com/jdxyw/generativeart"
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util"
	"math/rand"
	"time"

	"github.com/jdxyw/generativeart/arts"
)

func Square(cfg *models.Config) string {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(cfg.Canvas.Width, cfg.Canvas.Height)
	c.SetBackground(util.RnColor())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())

	c.Draw(arts.NewRandomShape(util.Rn(130, 170)))

	cfg.Out.Prefix = "square"
	url := cfg.OutURL()
	c.ToPNG(url)

	return url
}
