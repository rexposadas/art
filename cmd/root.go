/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "art",
	Short: "A CLI for generating art",
	Long: `
	A CLI to easily generate art.


`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var file string // Files will be needed for a lot of commands.  So, it's at the top level.
var count string
var output string // where the output files will go. defaults to "output"

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "file to process")
	rootCmd.PersistentFlags().StringVarP(&output, "output", "o", "", "output directory")
	rootCmd.PersistentFlags().StringVarP(&count, "count", "c", "1", "count")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
