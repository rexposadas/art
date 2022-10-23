package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"image/color"
	"strings"
)

var squaresSpiralCmd = &cobra.Command{
	Use:   "spiral",
	Short: "squares making a spiral",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			out := samples.SpiralSquare(cfg)
			singWithText(out)
		}
	},
}

func singWithText(out string) {
	org := util.PathToImage(out)
	util.SignWithText(org, color.White, out)
	out = strings.Replace(out, ".png", "_black.png", 1)
	util.SignWithText(org, color.Black, out)
}

func sign(out, sig string) {
	pic := util.PathToImage(out)
	signature := util.PathToImage(sig)
	util.Sign(pic, signature)
}

func init() {
	squaresCmd.AddCommand(squaresSpiralCmd)
}
