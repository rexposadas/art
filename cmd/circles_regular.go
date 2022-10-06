package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"sync"
)

// circlesRegularCmd represents the circleregular command
var circlesRegularCmd = &cobra.Command{
	Use:   "regular",
	Short: "regular",
	Long:  `regular circles`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				samples.Circles(cfg)
			}()
		}

		wg.Wait()
	},
}

func init() {
	circlesCmd.AddCommand(circlesRegularCmd)
}
