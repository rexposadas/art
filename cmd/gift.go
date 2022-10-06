package cmd

import (
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/csv"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
)

// giftCmd represents the gift command
var giftCmd = &cobra.Command{
	Use:   "gift",
	Short: "runs files through gift library",
	Long: `

./art gift -f list.csv

list.csv is a list of file locations:
testdata/first.png
testdata/second.png

NOTE: only takes png files. 


`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		list := csv.ParseCSV(file)

		for _, line := range list {
			util.Gift(line[0])
		}
	},
}

func init() {
	rootCmd.AddCommand(giftCmd)
}
