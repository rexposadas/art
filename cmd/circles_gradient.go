package cmd

import (
	"github.com/rexposadas/art/util/config"
	"sync"

	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util/require"

	"github.com/spf13/cobra"
)

var circlesGradientCmd = &cobra.Command{
	Use:   "gradient",
	Short: "circles with gradient",
	Long:  `circles with gradient`,
	Run: func(cmd *cobra.Command, args []string) {

		require.FileName(file)
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := config.New(file)
		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				models.CircleGradient(cfg)

			}()
		}

		wg.Wait()
	},
}

func init() {
	circlesCmd.AddCommand(circlesGradientCmd)
}
