package aes

import (
	"log"
	"testing"

	"github.com/gzylg/kits/random"
)

var (
	key     = random.Str(16)
	content = random.Str(32)
)

//! ===================== CBC =====================

func TestCBC(t *testing.T) {
	log.Println("content:", content)

	a, err := CBCEncrypt([]byte(content), []byte(key))
	log.Println("Encrypt", a, err)

	b, err := CBCDecrypt(a, []byte(key))
	log.Println("Decrypt", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}

func TestCBCBase64(t *testing.T) {
	log.Println("content:", content)

	a, err := CBCEncryptToBase64(content, key)
	log.Println("EncryptToBase64", a, err)

	b, err := CBCDecryptFromBase64(a, key)
	log.Println("DecryptFromBase64", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}

//! ===================== ECB =====================

func TestECB(t *testing.T) {
	log.Println("content:", content)

	a, err := ECBEncrypt([]byte(content), []byte(key))
	log.Println("Encrypt", a, err)

	b, err := ECBDecrypt(a, []byte(key))
	log.Println("Decrypt", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}

func TestEBCBase64(t *testing.T) {
	log.Println("content:", content)

	a, err := ECBEncryptToBase64(content, key)
	log.Println("EncryptToBase64", a, err)

	b, err := ECBDecryptFromBase64(a, key)
	log.Println("DecryptFromBase64", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}

//! ===================== CFB =====================

func TestCBF(t *testing.T) {
	log.Println("content:", content)

	a, err := CFBEncrypt([]byte(content), []byte(key))
	log.Println("Encrypt", a, err)

	b, err := CFBDecrypt(a, []byte(key))
	log.Println("Decrypt", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}

func TestCBFBase64(t *testing.T) {
	log.Println("content:", content)

	a, err := CFBEncryptToBase64(content, key)
	log.Println("EncryptToBase64", a, err)

	b, err := CFBDecryptFromBase64(a, key)
	log.Println("DecryptFromBase64", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}

//! ===================== CTR =====================

func TestCTR(t *testing.T) {
	log.Println("content:", content)

	a, err := CTREncrypt([]byte(content), []byte(key))
	log.Println("Encrypt", a, err)

	b, err := CTRDecrypt(a, []byte(key))
	log.Println("Decrypt", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}

func TestCTRBase64(t *testing.T) {
	log.Println("content:", content)

	a, err := CTREncryptToBase64(content, key)
	log.Println("EncryptToBase64", a, err)

	b, err := CTRDecryptFromBase64(a, key)
	log.Println("DecryptFromBase64", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}
