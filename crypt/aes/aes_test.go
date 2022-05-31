package aes

import (
	"log"
	"testing"
	"time"

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

//! ===================== 速度测试 =====================
func TestRunTime(t *testing.T) {
	count := 10000 // 运行次数

	//* -------------- CBC
	start := time.Now()
	for i := 0; i < count; i++ {
		c := random.Str(128)
		k := random.Str(16)

		a, _ := CBCEncryptToBase64(c, k)
		b, _ := CBCDecryptFromBase64(a, k)

		if string(b) != c {
			t.Fatal("crypt Failed.")
			return
		}
	}
	log.Println("CBC:", time.Since(start))

	//* -------------- ECB
	start = time.Now()
	for i := 0; i < count; i++ {
		c := random.Str(128)
		k := random.Str(16)

		a, _ := ECBEncryptToBase64(c, k)
		b, _ := ECBDecryptFromBase64(a, k)

		if string(b) != c {
			t.Fatal("crypt Failed.")
			return
		}
	}
	log.Println("ECB:", time.Since(start))

	//* -------------- CFB
	start = time.Now()
	for i := 0; i < count; i++ {
		c := random.Str(128)
		k := random.Str(16)

		a, _ := CFBEncryptToBase64(c, k)
		b, _ := CFBDecryptFromBase64(a, k)

		if string(b) != c {
			t.Fatal("crypt Failed.")
			return
		}
	}
	log.Println("CFB:", time.Since(start))

	//* -------------- CTR
	start = time.Now()
	for i := 0; i < count; i++ {
		c := random.Str(128)
		k := random.Str(16)

		a, _ := CTREncryptToBase64(c, k)
		b, _ := CTRDecryptFromBase64(a, k)

		if string(b) != c {
			t.Fatal("crypt Failed.")
			return
		}
	}
	log.Println("CTR:", time.Since(start))
}
