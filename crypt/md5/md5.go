package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func Gen(content []byte) (encrypted string, err error) {
	h := md5.New()
	_, err = h.Write(content)

	return hex.EncodeToString(h.Sum(nil)), err
}
