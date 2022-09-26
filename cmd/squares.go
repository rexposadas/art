package cmd

import (
	"github.com/spf13/cobra"
)

// squaresCmd represents the square command
var squaresCmd = &cobra.Command{
	Use:   "squares",
	Short: "generate square art",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(squaresCmd)
}
