/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"art/util"
	"art/util/csv"
	"art/util/require"
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
