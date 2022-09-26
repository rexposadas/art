package cmd

import (
	"github.com/spf13/cobra"
	"log"
)

// circlesCmd represents the circles command
var circlesCmd = &cobra.Command{
	Use:   "circles",
	Short: "circles",
	Long:  `circle commands`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("look into the sub commands for actual generated art")
	},
}

func init() {
	rootCmd.AddCommand(circlesCmd)
}
