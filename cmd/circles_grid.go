package cmd

import (
	"sync"

	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
)

var circlesGridCmd = &cobra.Command{
	Use:   "grid",
	Short: "circles with grid",
	Long: `
Circles with grid.
`,
	Run: func(cmd *cobra.Command, args []string) {
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := config.New(file)
		for i := 0; i < total; i++ {
			wg.Add(1)

			go func() {
				defer wg.Done()
				models.CirclesGrid(cfg)

			}()

		}

		wg.Wait()
	},
}

func init() {
	circlesCmd.AddCommand(circlesGridCmd)
}
