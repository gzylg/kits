package pkcs5

import (
	"bytes"

	"github.com/gzylg/kits/errs"
)

func Padding(plainText []byte, blockSize int) []byte {
	padding := blockSize - (len(plainText) % blockSize)
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	newText := append(plainText, padText...)
	return newText
}

func UnPadding(plainText []byte) ([]byte, error) {
	length := len(plainText)
	number := int(plainText[length-1])
	if number >= length {
		return nil, errs.New("padding size error: please check the secret key.")
	}
	return plainText[:length-number], nil
}
