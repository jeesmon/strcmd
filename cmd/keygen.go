package cmd

import (
	"fmt"

	"github.com/coderanger/controller-utils/randstring"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(keygenCmd)
}

var keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Generate Encrypt Key",
	Run: func(cmd *cobra.Command, args []string) {
		keygen(cmd, args)
	},
}

func keygen(cmd *cobra.Command, args []string) {
	fmt.Println(randstring.MustRandomString(32)[:32])
}
