package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"sync"

	"github.com/spf13/cobra"
)

var smokeCmd = &cobra.Command{
	Use:   "smoke",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := config.New(file)
		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				models.Smoke(cfg)
			}()

		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(smokeCmd)
}
