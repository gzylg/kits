package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"log"
	"runtime"

	"github.com/gzylg/kits/crypt/pkcs5"
)

/* Cipher Block Chaining mode（分组密码链接模式） */

func CBCEncrypt(plainText, key []byte) (cipherText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize() // 获取秘钥块的长度

	paddingText := pkcs5.Padding(plainText, blockSize) // 补全码

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式

	cipherText = make([]byte, len(paddingText)) // 创建数组

	blockMode.CryptBlocks(cipherText, paddingText) // 加密
	return
}

func CBCEncryptToBase64(plainText, key string) (cipherText string, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", ErrKeyLengthSixteen
	}

	var b []byte
	if b, err = CBCEncrypt([]byte(plainText), []byte(key)); err != nil {
		return
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func CBCDecrypt(cipherText, key []byte) (plainText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	block, err := aes.NewCipher(key) // 分组秘钥
	if err != nil {
		return nil, err
	}

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key or text is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()

	blockSize := block.BlockSize() // 获取秘钥块的长度

	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式

	paddingText := make([]byte, len(cipherText)) // 创建数组

	blockMode.CryptBlocks(paddingText, cipherText) // 解密

	plainText, err = pkcs5.UnPadding(paddingText) // 去除补全码

	return
}

func CBCDecryptFromBase64(cipherText, key string) (plainText []byte, err error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return nil, ErrKeyLengthSixteen
	}

	var b []byte
	if b, err = base64.StdEncoding.DecodeString(cipherText); err != nil {
		return
	}

	return CBCDecrypt(b, []byte(key))
}
