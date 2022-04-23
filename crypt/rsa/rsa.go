package rsa

import (
	"bytes"
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"log"
	"runtime"

	"github.com/gzylg/kits/file"
)

func init() {
	log.SetFlags(log.Ldate | log.Lshortfile)
}

const (
	privateKeyPrefix = " Y-Server RSA PRIVATE KEY "
	publicKeyPrefix  = " Y-Server RSA PUBLIC KEY "
)

// GenRsaKey 生成rsa秘钥对
func GenRsaKey(bits int, publicKeyFile, privateKeyFile string) (p *[]byte, err error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, err
	}
	x509PrivateKey := x509.MarshalPKCS1PrivateKey(privateKey)
	//privateFile, err := os.Create(savePath + privateFileName)
	//if err != nil {
	//	return nil, err
	//}
	//defer privateFile.Close()
	privateBlock := pem.Block{
		Type:  privateKeyPrefix,
		Bytes: x509PrivateKey,
	}

	if err = file.WriteByte(privateKeyFile, pem.EncodeToMemory(&privateBlock)); err != nil {
		return nil, err
	}
	//if err = pem.Encode(privateFile, &privateBlock); err != nil {
	//	return nil, err
	//}
	publicKey := privateKey.PublicKey
	x509PublicKey, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil, err
	}
	//publicFile, _ := os.Create(savePath + publicFileName)
	//defer publicFile.Close()
	publicBlock := pem.Block{
		Type:  publicKeyPrefix,
		Bytes: x509PublicKey,
	}
	rep := pem.EncodeToMemory(&publicBlock)
	if err = file.WriteByte(publicKeyFile, rep); err != nil {
		return nil, err
	}
	//if err = pem.Encode(publicFile, &publicBlock); err != nil {
	//	return nil, err
	//}

	return &rep, nil
}

// VerifyPubKey 校验公钥
func VerifyPubKey(key []byte) bool {
	block, _ := pem.Decode(key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	if _, err := x509.ParsePKIXPublicKey(block.Bytes); err != nil {
		return false
	}
	return true
}

func Encrypt(plainText, key []byte) (cryptText []byte, err error) {
	block, _ := pem.Decode(key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()

	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey := publicKeyInterface.(*rsa.PublicKey)

	cipherText, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, plainText)
	if err != nil {
		return nil, err
	}
	return cipherText, nil
}

// EncryptBlock 使用公钥加密
func EncryptBlock(plainText, publicKeyByte []byte) (bytesEncrypt []byte, err error) {
	block, _ := pem.Decode(publicKeyByte)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}

	keySize, srcSize := publicKey.(*rsa.PublicKey).Size(), len(plainText)
	pub := publicKey.(*rsa.PublicKey)
	//logs.Debug("密钥长度：", keySize, "\t明文长度：\t", srcSize)
	//单次加密的长度需要减掉padding的长度，PKCS1为11
	offSet, once := 0, keySize-11
	buffer := bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + once
		if endIndex > srcSize {
			endIndex = srcSize
		}
		// 加密一部分
		bytesOnce, err := rsa.EncryptPKCS1v15(rand.Reader, pub, plainText[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesEncrypt = buffer.Bytes()
	return
}

func Decrypt(cryptText, key []byte) (plainText []byte, err error) {
	block, _ := pem.Decode(key)

	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return []byte{}, err
	}
	plainText, err = rsa.DecryptPKCS1v15(rand.Reader, privateKey, cryptText)
	if err != nil {
		return []byte{}, err
	}
	return plainText, nil
}

// DecryptBlock 使用私钥解密
func DecryptBlock(src, privateKeyByte []byte) (bytesDecrypt []byte, err error) {
	block, _ := pem.Decode(privateKeyByte)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return
	}

	//private := privateKey.(*rsa.PrivateKey)
	private := privateKey
	keySize, srcSize := private.Size(), len(src)
	//logs.Debug("密钥长度：", keySize, "\t密文长度：\t", srcSize)
	var offSet = 0
	var buffer = bytes.Buffer{}
	for offSet < srcSize {
		endIndex := offSet + keySize
		if endIndex > srcSize {
			endIndex = srcSize
		}
		bytesOnce, err := rsa.DecryptPKCS1v15(rand.Reader, private, src[offSet:endIndex])
		if err != nil {
			return nil, err
		}
		buffer.Write(bytesOnce)
		offSet = endIndex
	}
	bytesDecrypt = buffer.Bytes()
	return
}

// Sign 签名
func Sign(msg, Key []byte) (cryptText []byte, err error) {
	block, _ := pem.Decode(Key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	myHash := sha256.New()
	myHash.Write(msg)
	hashed := myHash.Sum(nil)
	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return nil, err
	}
	return sign, nil
}

// VerifySign 验签
func VerifySign(msg []byte, sign []byte, Key []byte) bool {
	block, _ := pem.Decode(Key)
	defer func() {
		if err := recover(); err != nil {
			switch err.(type) {
			case runtime.Error:
				log.Println("runtime err:", err, "Check that the key is correct")
			default:
				log.Println("error:", err)
			}
		}
	}()
	publicInterface, _ := x509.ParsePKIXPublicKey(block.Bytes)
	publicKey := publicInterface.(*rsa.PublicKey)
	myHash := sha256.New()
	myHash.Write(msg)
	hashed := myHash.Sum(nil)
	result := rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashed, sign)
	return result == nil
}
