package des

import (
	"log"
	"testing"

	"github.com/gzylg/kits/random"
)

var (
	key     = random.Str(8)
	content = random.Str(32)
)

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
