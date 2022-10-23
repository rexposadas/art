package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"sync"
)

// perlsCmd represents the perls command
var perlsCmd = &cobra.Command{
	Use:   "perls",
	Short: "Generate Perls",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		cfg := models.NewConfig(file)

		perls(cfg, total)
	},
}

func perls(cfg *models.Config, count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			out := samples.Perls(cfg)
			singWithText(out)
		}()
	}

	wg.Wait()

}

func init() {
	rootCmd.AddCommand(perlsCmd)
}
