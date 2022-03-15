package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	decCmd.Flags().StringVar(&encKey, "key", os.Getenv("STRCMD_ENCRYPT_KEY"), "Hex encoded encryption key")
	decCmd.Flags().StringVar(&input, "in", "", "Input string")
	decCmd.MarkFlagRequired("key")
	decCmd.MarkFlagRequired("in")
	rootCmd.AddCommand(decCmd)
}

var decCmd = &cobra.Command{
	Use:   "dec",
	Short: "Decrypt string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(dec(encKey, input))
	},
}

// Copy from https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
func dec(encKey, input string) string {
	key, err := hex.DecodeString(encKey)
	if err != nil {
		panic(err.Error())
	}

	enc, err := hex.DecodeString(input)
	if err != nil {
		panic(err.Error())
	}

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Get the nonce size
	nonceSize := aesGCM.NonceSize()

	//Extract the nonce from the encrypted data
	nonce, ciphertext := enc[:nonceSize], enc[nonceSize:]

	//Decrypt the data
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%s", plaintext)
}
