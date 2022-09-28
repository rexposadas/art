/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"art/samples"
	"fmt"

	"github.com/spf13/cobra"
)

// squaresRegularCmd represents the squaresregular command
var squaresRegularCmd = &cobra.Command{
	Use:   "regular",
	Short: "create basic square images",
	Long: `


`,
	Run: func(cmd *cobra.Command, args []string) {
		for i := 0; i < 20; i++ {
			samples.Square(fmt.Sprintf("square-%d", i))
		}

	},
}

func init() {
	squaresCmd.AddCommand(squaresRegularCmd)
}
