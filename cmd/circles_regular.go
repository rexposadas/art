package cmd

import (
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"sync"
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

func circlesRegular(cfg *config.Config, count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			out := samples.Circles(cfg)
			signWithText(out)
		}()
	}

	wg.Wait()
}

func init() {
	circlesCmd.AddCommand(circlesRegularCmd)
}
