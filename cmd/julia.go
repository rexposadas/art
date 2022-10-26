package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"sync"
)

// juliaCmd represents the julia command
var juliaCmd = &cobra.Command{
	Use:   "julia",
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
				models.Julia(cfg)
			}()
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(juliaCmd)
}
