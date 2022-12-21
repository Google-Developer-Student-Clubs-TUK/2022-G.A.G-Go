package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	b64 "encoding/base64"
	"fmt"
)

func AESEncrypt(src string, key []byte, iv []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
	}
	if src == "" {
		fmt.Println("plain content empty")
	}
	cbc := cipher.NewCBCEncrypter(block, iv)
	content := []byte(src)
	content = PKCS5Padding(content, block.BlockSize())
	crypted := make([]byte, len(content))
	cbc.CryptBlocks(crypted, content)

	return crypted
}

func AESDecrypt(crypt string, key []byte, iv []byte) (string, error) {
	cryptData, err := b64.StdEncoding.DecodeString(crypt)
	if err != nil {
		fmt.Println("key error1", err)
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("key error1", err)
		return "", err
	}

	if len(cryptData) == 0 {
		fmt.Println("plain content empty")
		return "", err
	}

	cbc := cipher.NewCBCDecrypter(block, []byte(iv))
	decrypted := make([]byte, len(cryptData))
	cbc.CryptBlocks(decrypted, cryptData)

	return string(PKCS5Trimming(decrypted)), nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5Trimming(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
