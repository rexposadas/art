package cmd

import (
	"art/models"
	"art/samples"
	"art/util/require"
	"github.com/spf13/cobra"
)

// circlesCmd represents the circles command
var circlesCmd = &cobra.Command{
	Use:   "circles",
	Short: "generates s",
	Long: `

given input parameters, generate circles 

./art circles -c 5 -f input/s.json

`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		total := require.Count(count)

		cfg := models.NewConfig(file)
		for i := 0; i < total; i++ {
			samples.Circles(cfg)
		}
	},
}

func init() {
	rootCmd.AddCommand(circlesCmd)

}
