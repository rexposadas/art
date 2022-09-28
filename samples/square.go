package samples

import (
	"art/util"
	"fmt"
	"github.com/jdxyw/generativeart"
	"github.com/jdxyw/generativeart/arts"
	"math/rand"
	"time"
)

func Square(filename string) {
	rand.Seed(time.Now().Unix())
	c := generativeart.NewCanva(1080, 1080)
	c.SetBackground(util.RnColor())
	c.FillBackground()
	c.SetColorSchema(util.RnColorScheme())
	//c.Draw(arts.NewRandomShape(150))
	c.Draw(arts.NewRandomShape(util.Rn(130, 170)))
	c.ToPNG(fmt.Sprintf("output/%s.png", filename))
}
