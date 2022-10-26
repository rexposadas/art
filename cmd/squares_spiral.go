package cmd

import (
	"github.com/jdxyw/generativeart"
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"image/color"
	"log"
)

var squaresSpiralCmd = &cobra.Command{
	Use:   "spiral",
	Short: "squares making a spiral",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		cfg := config.New(file)
		for i := 0; i < total; i++ {
			out := models.SpiralSquare(cfg)
			signWithText(out)

		}
	},
}

func not255(o uint8) uint8 {

	if o == 255 {
		return 0
	}

	x := o + 10
	if x == 255 {
		return x - 10
	}

	return x

}

func changeColor(orig color.Color) color.Color {
	//c1 := color.Gray16Model.Convert(orig).(color.Gray16)
	result := color.RGBAModel.Convert(orig).(color.RGBA)

	return color.RGBA{
		R: result.G,
		G: result.B,
		B: result.R,
		A: result.A,
	}

	//return result
}

func signWithText(out string) {
	org := util.PathToImage(out)

	newCanva := generativeart.NewCanva(1080, 1080)
	newImg := newCanva.Img()

	for x := 0; x < org.Bounds().Max.X; x++ {
		for y := 0; y < org.Bounds().Max.Y; y++ {
			at := org.At(x, y)
			newAT := changeColor(at)
			newImg.Set(x, y, newAT)
		}
	}

	if err := newCanva.ToPNG("output/reworked.png"); err != nil {
		log.Printf("writing jpeg %s", err)
	}

	util.SignWithText(org, color.White, out)
	//out = strings.Replace(out, ".png", "_black.png", 1)
	//util.SignWithText(org, color.Black, out)
}

func sign(out, sig string) {
	pic := util.PathToImage(out)
	signature := util.PathToImage(sig)
	util.Sign(pic, signature)
}

func init() {
	squaresCmd.AddCommand(squaresSpiralCmd)
}
