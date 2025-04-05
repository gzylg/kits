package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

func CFBEncrypt(plainText []byte, key []byte) (cipherText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return nil, err
	}

	cipherText = make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	return
}

func CFBEncryptToBase64(plainText string, key string) (cipherText string, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", ErrKeyLengthSixteen
	}

	var b []byte
	if b, err = CFBEncrypt([]byte(plainText), []byte(key)); err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func CFBDecrypt(cipherText []byte, key []byte) (plainText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return nil, err
	}

	if len(cipherText) < aes.BlockSize {
		return nil, err
	}
	iv := cipherText[:aes.BlockSize]
	plainText = cipherText[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plainText, plainText)
	return
}

func CFBDecryptFromBase64(cipherText string, key string) (plainText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	var b []byte
	if b, err = base64.StdEncoding.DecodeString(cipherText); err != nil {
		return
	}

	return CFBDecrypt(b, []byte(key))
}
