package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
)

// squaresspiralCmd represents the squaresspiral command
var squaresspiralCmd = &cobra.Command{
	Use:   "spiral",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			samples.SpiralSquare(cfg)
		}
	},
}

func init() {
	squaresCmd.AddCommand(squaresspiralCmd)
}
