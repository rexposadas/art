package cmd

import (
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"sync"

	"github.com/spf13/cobra"
)

// squaresRegularCmd represents the squaresregular command
var squaresRegularCmd = &cobra.Command{
	Use:   "regular",
	Short: "create basic square images",
	Long: `


`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := config.New(file)
		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				out := samples.Square(cfg)
				singWithText(out)
			}()

		}

		wg.Wait()
	},
}

func init() {
	squaresCmd.AddCommand(squaresRegularCmd)
}
