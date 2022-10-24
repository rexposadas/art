package cmd

import (
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util/config"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"sync"
)

// juliaCmd represents the julia command
var juliaCmd = &cobra.Command{
	Use:   "julia",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		var wg sync.WaitGroup

		cfg := config.New(file)
		for i := 0; i < total; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				samples.Julia(cfg)
			}()
		}

		wg.Wait()
	},
}

func init() {
	rootCmd.AddCommand(juliaCmd)
}
