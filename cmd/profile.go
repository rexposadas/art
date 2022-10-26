/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"github.com/jdxyw/generativeart"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"log"
)

// profileCmd represents the profile command
var profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "takes a profile pic or any pic and changes colors",
	Long: `
Meant to take any pic and have fun with it .

`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)

		changeProfile(file)

	},
}

func changeProfile(out string) {
	org := util.PathToImageJPG(out)

	newCanva := generativeart.NewCanva(org.Bounds().Max.X, org.Bounds().Max.Y)

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

	//out = strings.Replace(out, ".png", "_black.png", 1)
	//util.SignWithText(org, color.Black, out)
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
