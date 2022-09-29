package cmd

import (
	"art/models"
	"art/samples"
	"art/util/require"
	"sync"

	"github.com/spf13/cobra"
)

// circlesGradientCmd .
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
				samples.CircleGradient(cfg)
			}()
		}

		wg.Wait()
	},
}

func init() {
	circlesCmd.AddCommand(circlesGradientCmd)
}
