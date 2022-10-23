package cmd

import (
	"fmt"
	"github.com/rexposadas/art/util"
	"github.com/rexposadas/art/util/require"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
)

var mergeCmd = &cobra.Command{
	Use:   "merge",
	Short: "automated way of merging what we have so far",
	Long: `

	Merge photos with the generated images

	./art merge -f input/list.csv

`,
	Run: func(cmd *cobra.Command, args []string) {
		require.FileName(file)
		util.Merge("new.png", "new.png")
	},
}

func fileList(look string) []string {
	files, err := ioutil.ReadDir(look)
	if err != nil {
		log.Fatal(err)
	}

	var list []string
	for _, file := range files {
		fmt.Println(file.Name(), file.IsDir())
		if !file.IsDir() {
			list = append(list, file.Name())
		}

	}

	log.Printf("returning %+v", list)

	return list
}

func init() {
	rootCmd.AddCommand(mergeCmd)
}
