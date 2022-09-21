package cmd

import (
	"github.com/spf13/cobra"
)

// tryCmd represents the try command
var tryCmd = &cobra.Command{
	Use:   "try",
	Short: "where we try stuff",
	Long: `

short term place where we try stuff

`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	rootCmd.AddCommand(tryCmd)
}
