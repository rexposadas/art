package cmd

import (
	"github.com/rexposadas/art/util/config"
	"github.com/spf13/cobra"
)

var circlesCmd = &cobra.Command{
	Use:   "circles",
	Short: "circles",
	Long:  `circle commands`,
	Run: func(cmd *cobra.Command, args []string) {

		// Default command to create circles.
		circlesRegular(config.Default(), 5)
	},
}

func init() {
	rootCmd.AddCommand(circlesCmd)
}
