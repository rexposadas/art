package cmd

import (
	"sync"

	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/samples"
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

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				out := samples.CircleGradient(cfg)
				singWithText(out)
			}()
		}

		wg.Wait()
	},
}

func init() {
	circlesCmd.AddCommand(circlesGradientCmd)
}
