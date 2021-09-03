package crowdeAesCbc256

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"log"
)

type AesEncrypt struct {
	Iv string
	Secret string
}

func (s *AesEncrypt) Init(iv string, secret string) {
	if iv == "" {
		log.Fatal(fmt.Printf("Error: %s \n", "initial vector is missing"))
	}
	if len(iv) != 16 {
		log.Fatal(fmt.Printf("Error: %s \n", "initial vector should be 16 digits"))
	}
	if secret == "" {
		log.Fatal(fmt.Printf("Error: %s \n", "secret is missing"))
	}
	s.Iv = iv
	s.Secret = secret
}

func (s *AesEncrypt) AESEncrypt(plaintext string) string {
	// block chiper
	block, err := aes.NewCipher([]byte(s.Secret))
	if err != nil {
		log.Fatal(fmt.Printf("Error: %#v \n", err))
	}
	if plaintext == "" {
		log.Fatal(fmt.Printf("Error: %s \n", "plaintext is empty"))
	}
	ecb := cipher.NewCBCEncrypter(block, []byte(s.Iv))
	content := PKCS5Padding([]byte(plaintext), block.BlockSize())

	crypted := make([]byte, len(content))
	ecb.CryptBlocks(crypted, content)

	// return as base 64
	return base64.StdEncoding.EncodeToString(crypted)
}

func (s *AesEncrypt) AESDecrypt(encrypted string) []byte {
	// decode encrypted data
	encryptedData, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		log.Fatal(fmt.Printf("Error: %s \n", "unable to decode encrypted data"))
	}
	// block chiper
	block, err := aes.NewCipher([]byte(s.Secret))
	if err != nil {
		log.Fatal(fmt.Printf("Error: %#v \n", err))
	}
	if len(encryptedData) == 0 {
		log.Fatal(fmt.Printf("Error: %s \n", "encrypted data is empty"))
	}
	ecb := cipher.NewCBCDecrypter(block, []byte(s.Iv))

	decrypted := make([]byte, len(encryptedData))
	ecb.CryptBlocks(decrypted, encryptedData)

	// return as string
	return PKCS5Trimming(decrypted)
}