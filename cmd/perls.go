package cmd

import (
	"github.com/rexposadas/art/models"
	"github.com/rexposadas/art/samples"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
)

// perlsCmd represents the perls command
var perlsCmd = &cobra.Command{
	Use:   "perls",
	Short: "Generate Perls",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			samples.Perls(cfg)
		}
	},
}

func init() {
	rootCmd.AddCommand(perlsCmd)
}
