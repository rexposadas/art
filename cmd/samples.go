package cmd

import (
	"art/samples"
	"fmt"
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
		generateSamples()
	},
}

func generateSamples() {
	for i := 0; i < 3; i++ {
		//samples.Circles(fmt.Sprintf("s-%d", i))
		samples.Square(fmt.Sprintf("square-%d", i))
		samples.CircleGradient(fmt.Sprintf("gradient-%d", i))
	}
}

func init() {
	rootCmd.AddCommand(samplesCmd)
}
