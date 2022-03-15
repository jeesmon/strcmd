package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	encCmd.Flags().StringVar(&encKey, "key", os.Getenv("STRCMD_ENCRYPT_KEY"), "Hex encoded encryption key")
	encCmd.Flags().StringVar(&input, "in", "", "Input string")
	encCmd.MarkFlagRequired("key")
	encCmd.MarkFlagRequired("in")
	rootCmd.AddCommand(encCmd)
}

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "Encrypt string",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(enc(encKey, input))
	},
}

// Copy from https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
func enc(encKey, input string) string {
	key, err := hex.DecodeString(encKey)
	if err != nil {
		panic(err.Error())
	}

	plaintext := []byte(input)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}
