package cmd

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	encCmd.Flags().StringVar(&encKey, "key", os.Getenv("STRCMD_ENCRYPT_KEY"), "Encryption Key")
	encCmd.Flags().StringVar(&input, "in", "", "Input String")
	encCmd.MarkFlagRequired("key")
	encCmd.MarkFlagRequired("in")
	rootCmd.AddCommand(encCmd)
}

var encCmd = &cobra.Command{
	Use:   "enc",
	Short: "Encrypt String",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(enc(encKey, input))
	},
}

// Copy from https://www.melvinvivas.com/how-to-encrypt-and-decrypt-data-using-aes
func enc(encKey, input string) string {
	//Since the key is in string, we need to convert it to bytes
	key := []byte(fmt.Sprintf("%32v", encKey))
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
