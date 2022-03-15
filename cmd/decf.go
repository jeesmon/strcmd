package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	decfCmd.Flags().StringVar(&encKey, "key", os.Getenv("STRCMD_ENCRYPT_KEY"), "Hex encoded encryption key")
	decfCmd.Flags().StringVar(&file, "file", "", "Input file")
	decfCmd.MarkFlagRequired("key")
	decfCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(decfCmd)
}

var decfCmd = &cobra.Command{
	Use:   "decf",
	Short: "Decrypt strings from a file",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err.Error())
		}

		content := string(b)
		lines := strings.Split(content, "\n")
		for _, l := range lines {
			if l != "" {
				fmt.Println(dec(encKey, l))
			}
		}
	},
}
