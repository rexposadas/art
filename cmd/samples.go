package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/config"
	"github.com/spf13/cobra"
	"log"
)

// samplesCmd represents the models command
var samplesCmd = &cobra.Command{
	Use:   "models",
	Short: "generate models",
	Long: `
		generate the models.

Usage: 
	./art models
`,
	Run: func(cmd *cobra.Command, args []string) {
		if err := util.MakeDefaultDir(); err != nil {
			log.Printf("failed to created output directory")
			return
		}

		models.Circles(config.Default())
		models.SpiralSquare(config.Default())
	},
}

func init() {
	rootCmd.AddCommand(samplesCmd)
}
