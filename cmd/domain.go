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

// domainCmd represents the domain command
var domainCmd = &cobra.Command{
	Use:   "domain",
	Short: "domain",
	Long: `
domain

`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				samples.Domain(cfg)
			}()
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(domainCmd)
}
