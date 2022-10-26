package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"sync"
)

var perlsCmd = &cobra.Command{
	Use:   "perls",
	Short: "Generate Perls",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		total := require.Count(count)
		cfg := config.New(file)

		perls(cfg, total)
	},
}

func perls(cfg config.Config, count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			models.Perls(cfg)
		}()
	}

	wg.Wait()

}

func init() {
	rootCmd.AddCommand(perlsCmd)
}
