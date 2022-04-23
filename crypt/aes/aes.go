package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
)

// =================== CBC ======================
func EncryptCBCBase64(origData, key string) (err error, encrypted string) {
	var b []byte
	if err, b = EncryptCBC([]byte(origData), []byte(key)); err != nil {
		return err, ""
	}
	return nil, base64.StdEncoding.EncodeToString(b)
}
func DecryptCBCFromBase64(encrypted string, key string) (err error, decrypted string) {
	var b []byte
	if b, err = base64.StdEncoding.DecodeString(encrypted); err != nil {
		return err, ""
	}

	if err, b = DecryptCBC(b, []byte(key)); err != nil {
		return err, ""
	}
	return nil, string(b)
}
func EncryptCBC(origData, key []byte) (err error, encrypted []byte) {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, err := aes.NewCipher(key)
	if err != nil {
		return err, nil
	}
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	origData = pkcs5Padding(origData, blockSize)                // 补全码
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize]) // 加密模式
	encrypted = make([]byte, len(origData))                     // 创建数组
	blockMode.CryptBlocks(encrypted, origData)                  // 加密
	return nil, encrypted
}
func DecryptCBC(encrypted, key []byte) (err error, decrypted []byte) {
	block, err := aes.NewCipher(key) // 分组秘钥
	if err != nil {
		return err, nil
	}
	blockSize := block.BlockSize()                              // 获取秘钥块的长度
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize]) // 加密模式
	decrypted = make([]byte, len(encrypted))                    // 创建数组
	blockMode.CryptBlocks(decrypted, encrypted)                 // 解密
	decrypted = pkcs5UnPadding(decrypted)                       // 去除补全码
	return nil, decrypted
}
func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// =================== ECB(速度应该是最快的) ======================
func EncryptECBBase64(origData, key string) (err error, encrypted string) {
	var b []byte
	if err, b = EncryptECB([]byte(origData), []byte(key)); err != nil {
		return err, ""
	}
	return nil, base64.StdEncoding.EncodeToString(b)
}
func DecryptECBFromBase64(encrypted string, key string) (err error, decrypted string) {
	var b []byte
	if b, err = base64.StdEncoding.DecodeString(encrypted); err != nil {
		return err, ""
	}
	if err, b = DecryptECB(b, []byte(key)); err != nil {
		return err, ""
	}
	return nil, string(b)
}
func EncryptECB(origData, key []byte) (err error, encrypted []byte) {
	cipher, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return err, nil
	}
	length := (len(origData) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, origData)
	pad := byte(len(plain) - len(origData))
	for i := len(origData); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cipher.BlockSize(); bs <= len(origData); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return nil, encrypted
}
func DecryptECB(encrypted, key []byte) (err error, decrypted []byte) {
	cipher, err := aes.NewCipher(generateKey(key))
	if err != nil {
		return err, nil
	}
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cipher.BlockSize(); bs < len(encrypted); bs, be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return nil, decrypted[:trim]
}
func generateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// =================== CFB（速度最慢的） ======================
func EncryptCFB(origData, key []byte) (err error, encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err, nil
	}
	if err != nil {
		panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return nil, encrypted
}
func DecryptCFB(encrypted, key []byte) (err error, decrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err, nil
	}
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return nil, encrypted
}
