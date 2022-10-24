package cmd

import (
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"sync"
)

var circlesGridCmd = &cobra.Command{
	Use:   "grid",
	Short: "circles with grid",
	Long: `
Circles with grid.
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
				out := samples.CirclesGrid(cfg)
				singWithText(out)
			}()

		}

		wg.Wait()
	},
}

func init() {
	circlesCmd.AddCommand(circlesGridCmd)
}
