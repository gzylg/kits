package str

import (
	"log"
	"testing"
)

func TestGetStr(t *testing.T) {
	str := "123abc456DEF789+-*/"
	log.Println("获得字符串中的英文字母：", GetEngLetter(str))
	log.Println("获得字符串中的大写字母：", GetCapital(str))
	log.Println("获得字符串中的小写字母：", GetLower(str))
	log.Println("获得字符串中的数字：", GetNum(str))
}
