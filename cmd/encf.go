package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	encfCmd.Flags().StringVar(&encKey, "key", os.Getenv("STRCMD_ENCRYPT_KEY"), "Encryption Key")
	encfCmd.Flags().StringVar(&file, "file", "", "Input File")
	encfCmd.MarkFlagRequired("key")
	encfCmd.MarkFlagRequired("file")
	rootCmd.AddCommand(encfCmd)
}

var encfCmd = &cobra.Command{
	Use:   "encf",
	Short: "Encrypt String from a file",
	Run: func(cmd *cobra.Command, args []string) {
		b, err := ioutil.ReadFile(file)
		if err != nil {
			panic(err.Error())
		}

		content := string(b)
		lines := strings.Split(content, "\n")
		for _, l := range lines {
			if l != "" {
				fmt.Println(enc(encKey, l))
			}
		}
	},
}
