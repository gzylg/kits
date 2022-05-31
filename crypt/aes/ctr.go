package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

func CTREncrypt(plainText []byte, key []byte) (cipherText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()                  // 获取秘钥块的长度
	stream := cipher.NewCTR(block, key[:blockSize]) // 加密模式

	cipherText = make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText, nil
}

func CTREncryptToBase64(plainText string, key string) (cipherText string, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", ErrKeyLengthSixteen
	}

	var b []byte
	if b, err = CTREncrypt([]byte(plainText), []byte(key)); err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func CTRDecrypt(cipherText []byte, key []byte) (plainText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize() // 获取秘钥块的长度
	stream := cipher.NewCTR(block, key[:blockSize])
	plainText = make([]byte, len(cipherText))
	stream.XORKeyStream(plainText, cipherText)

	return plainText, nil
}

func CTRDecryptFromBase64(cipherText string, key string) (plainText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	var b []byte
	if b, err = base64.StdEncoding.DecodeString(cipherText); err != nil {
		return
	}

	return CTRDecrypt(b, []byte(key))
}
