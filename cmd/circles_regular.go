package cmd

import (
	"github.com/spf13/cobra"
	"sync"

	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
)

var circlesRegularCmd = &cobra.Command{
	Use:   "regular",
	Short: "regular",
	Long:  `regular circles`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		cfg := config.New(file)
		circlesRegular(cfg, total)
	},
}

func circlesRegular(cfg config.Config, count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			models.Circles(cfg)
		}()
	}

	wg.Wait()
}

func init() {
	circlesCmd.AddCommand(circlesRegularCmd)
}
