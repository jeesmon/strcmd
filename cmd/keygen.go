package cmd

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	keygenCmd.Flags().IntVar(&keySize, "size", 32, "Key size")
	rootCmd.AddCommand(keygenCmd)
}

var keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Generate hex encoded encrypt key",
	Run: func(cmd *cobra.Command, args []string) {
		keygen(cmd, args)
	},
}

func keygen(cmd *cobra.Command, args []string) {
	bytes := make([]byte, keySize) //generate a random byte key
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	key := hex.EncodeToString(bytes) //encode key in bytes to string for saving
	fmt.Println(key)
}
