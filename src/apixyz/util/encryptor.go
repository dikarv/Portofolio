package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"

	"github.com/spf13/viper"
)

func EncryptAES(value string) string {
	c, err := aes.NewCipher([]byte(viper.GetString("salt")))
	if err != nil {
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("%x", gcm.Seal(nonce, nonce, []byte(value), nil))
}

func DecryptAES(value string) string {
	cipherText, err := hex.DecodeString(value)
	if err != nil {
		return ""
	}
	c, err := aes.NewCipher([]byte(viper.GetString("salt")))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return ""
	}

	nonce, ciphertext := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	return fmt.Sprintf("%s", plaintext)

}
