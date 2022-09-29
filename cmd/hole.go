/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"art/models"
	"art/samples"
	"art/util/require"
	"github.com/spf13/cobra"
	"sync"
)

// holeCmd represents the hole command
var holeCmd = &cobra.Command{
	Use:   "hole",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				samples.Hole(cfg)
			}()

		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(holeCmd)
}
