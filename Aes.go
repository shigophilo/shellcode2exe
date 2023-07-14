package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	zip "math/rand"
	"time"
)

func ctoAes(buf []byte) (string, []byte) {

	base64Str := base64.StdEncoding.EncodeToString(buf)
	key := RandStr(base64Str)
	encryptStr, err := Encrypt(key, []byte(base64Str))
	base64Str = base64.StdEncoding.EncodeToString(encryptStr)
	if err != nil {
		fmt.Println(err)
	}
	return base64Str, key
}
func RandStr(str string) []byte {
	r := zip.New(zip.NewSource(time.Now().Unix()))
	bytes := make([]byte, 16)
	for i := 0; i < 16; i++ {
		b := r.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return bytes
}

func Encrypt(key []byte, text []byte) ([]byte, error) {
	// Init Cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// Padding
	paddingLen := aes.BlockSize - (len(text) % aes.BlockSize)
	paddingText := bytes.Repeat([]byte{byte(paddingLen)}, paddingLen)
	textWithPadding := append(text, paddingText...)
	// Getting an IV
	ciphertext := make([]byte, aes.BlockSize+len(textWithPadding))
	iv := ciphertext[:aes.BlockSize]
	// Randomness
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	// Actual encryption
	cfbEncrypter := cipher.NewCFBEncrypter(block, iv)
	cfbEncrypter.XORKeyStream(ciphertext[aes.BlockSize:], textWithPadding)
	return ciphertext, nil
}
