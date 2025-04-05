package random

import (
	"crypto/rand"
	r "math/rand"
	"time"
)

var alphaStr = []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz`)

// randomCreateBytes 从 参数 alphabets 中随机抽取 参数n 数量的 bytes，并返回。 如果 参数 alphabets 为空，使用默认 alphaStr
func randomCreateBytes(n int, alphabets ...byte) []byte {
	if len(alphabets) == 0 {
		alphabets = alphaStr
	}
	var bytes = make([]byte, n)
	var randBy bool
	if num, err := rand.Read(bytes); num != n || err != nil {
		r.Seed(time.Now().UnixNano())
		randBy = true
	}
	for i, b := range bytes {
		if randBy {
			bytes[i] = alphabets[r.Intn(len(alphabets))]
		} else {
			bytes[i] = alphabets[b%byte(len(alphabets))]
		}
	}
	return bytes
}

// NumStr 返回随机获得0-9中数量为 参数n 的数字字符串
func NumStr(n int) string {
	alphaNum := []byte(`0123456789`)
	return string(randomCreateBytes(n, alphaNum...))
}

// Str 返回随机获得0-9 a-z A-Z中数量为 参数n 的字符串
func Str(n int) string {
	return string(randomCreateBytes(n))
}

// NumAndUpperStr 返回随机获得0-9 A-Z中数量为 参数n 的字符串
func NumAndUpperStr(n int) string {
	alphabets := []byte(`0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ`)
	return string(randomCreateBytes(n, alphabets...))
}

// Appoint 指定内容，指定长度
func Appoint(alphaStr string, n int) string {
	return string(randomCreateBytes(n, []byte(alphaStr)...))
}
