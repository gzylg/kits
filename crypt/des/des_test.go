package des

import (
	"encoding/base64"
	"log"
	"testing"
)

var (
	key    = "Skzd8+wlXNI+dj5lJG7rvT9U" // random.Str(8)
	ivbyte = []byte("azb_2019")

	//content = random.Str(32)
	content = `{"UID":"hy_57178310980076","timeSample":1700933329101,"token":"5457F61452610C95B836B65EB78B4D14E6BB5F036360847C"}`
)

func TestCBC(t *testing.T) {
	log.Println("content:", content)

	a, err := CBCEncrypt([]byte(content), []byte(key), ivbyte...)
	log.Println("Encrypt", a, err)
	log.Println("Encrypt", base64.StdEncoding.EncodeToString(a))

	b, err := CBCDecrypt(a, []byte(key), ivbyte...)
	log.Println("Decrypt", string(b), err)

	if string(b) != content {
		t.Fatal("crypt Failed.")
	}
}
