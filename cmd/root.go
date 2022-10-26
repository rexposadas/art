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

func init() {
	rootCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "file to process")
	rootCmd.PersistentFlags().StringVarP(&count, "count", "c", "1", "count")
}
