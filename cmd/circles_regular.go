/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"art/models"
	"art/samples"
	"art/util/require"
	"github.com/spf13/cobra"
)

// circlesRegularCmd represents the circleregular command
var circlesRegularCmd = &cobra.Command{
	Use:   "regular",
	Short: "regular",
	Long:  `regular circles`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			samples.Circles(cfg)
		}
	},
}

func init() {
	circlesCmd.AddCommand(circlesRegularCmd)
}
