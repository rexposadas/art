package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util"
	"github.com/spf13/cobra"
	"log"
)

// samplesCmd represents the samples command
var samplesCmd = &cobra.Command{
	Use:   "samples",
	Short: "generate samples",
	Long: `
		generate the samples.

Usage: 
	./art samples
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := util.MakeDefaultDir(); err != nil {
			log.Printf("failed to created output directory")
			return
		}

		samples.Circles(models.DefaultConfig())
		samples.SpiralSquare(models.DefaultConfig())
	},
}

func init() {
	rootCmd.AddCommand(samplesCmd)
}
