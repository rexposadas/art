package cmd

import (
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
	},
}

func init() {
	rootCmd.AddCommand(samplesCmd)
}
