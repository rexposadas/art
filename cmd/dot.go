package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"sync"

	"github.com/spf13/cobra"
)

// dotCmd represents the dot command
var dotCmd = &cobra.Command{
	Use:   "dot",
	Short: "generate dot images",
	Long: `
`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)

		total := require.Count(count)
		cfg := config.New(file)

		dot(cfg, total)
	},
}

func dot(cfg config.Config, count int) {
	var wg sync.WaitGroup
	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			models.Dot(cfg)
		}()
	}
	wg.Wait()
}

func init() {
	rootCmd.AddCommand(dotCmd)
}
