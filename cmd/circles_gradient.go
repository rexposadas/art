package cmd

import (
	"art/samples"
	"fmt"

	"github.com/spf13/cobra"
)

// circlesGradientCmd .
var circlesGradientCmd = &cobra.Command{
	Use:   "gradient",
	Short: "circles with gradient",
	Long:  `circles with gradient`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < 3; i++ {
			samples.CircleGradient(fmt.Sprintf("gradient-%d", i))
		}
	},
}

func init() {
	circlesCmd.AddCommand(circlesGradientCmd)
}
