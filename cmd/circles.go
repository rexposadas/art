package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/spf13/cobra"
)

// circlesCmd represents the circles command
var circlesCmd = &cobra.Command{
	Use:   "circles",
	Short: "circles",
	Long:  `circle commands`,
	Run: func(cmd *cobra.Command, args []string) {

		// Default command to create circles.
		circlesRegular(models.DefaultConfig(), 5)
	},
}

func init() {
	rootCmd.AddCommand(circlesCmd)
}
