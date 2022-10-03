package cmd

import (
	"art/models"
	"art/samples"
	"github.com/spf13/cobra"
)

// samplesCmd represents the samples command
var samplesCmd = &cobra.Command{
	Use:   "samples",
	Short: "generate samples",
	Long: `
generate the samples we use
`,
	Run: func(cmd *cobra.Command, args []string) {
		samples.Circles(models.DefaultConfig())
	},
}

func init() {
	rootCmd.AddCommand(samplesCmd)
}
