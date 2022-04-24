package str

import (
	"regexp"
	"strings"
)

// HideStrAuto 隐藏传入文本中的部分内容，自动判断类型为手机号或者邮箱等
func HideStrAuto(str string) string {
	pattern := `[a-zA-Z0-9]+@[a-zA-Z0-9\.]+\.[a-zA-Z0-9]+` // 判断邮箱
	if regexp.MustCompile(pattern).MatchString(str) {
		return HideEmail(str)
	}

	pattern = `^1[3|4|5|6|7|8|9][0-9]\d{8}$` // 判断手机号
	if regexp.MustCompile(pattern).MatchString(str) {
		return HidePhone(str)
	}

	return hideStr(str)
}

func HideEmail(email string) string {
	pattern := `([a-zA-Z0-9]+)@([a-zA-Z0-9\.]+\.[a-zA-Z0-9]+)`
	reg := regexp.MustCompile(pattern)
	submatch := reg.FindAllStringSubmatch(email, -1)
	if len(submatch) != 1 {
		return hideStr(email)
	}
	if len(submatch[0]) != 3 {
		return hideStr(email)
	}

	return hideStr(submatch[0][1]) + "@" + submatch[0][2]
}

func HidePhone(phone string) string {
	return phone[:3] + "****" + phone[len(phone)-4:]
}

func HideName(name string) string {
	words := ([]rune)(name)
	switch len(words) {
	case 2:
		return string(words[0]) + "*"
	case 3:
		return string(words[0]) + "*" + string(words[len(words)-1])
	case 4:
		return string(words[0]) + "**" + string(words[len(words)-1])
	}
	return name
}

// hideStr 隐藏传入文本的部分类容，仅对数字和字母有效
func hideStr(str string) string {
	var temp string
	switch len(str) {
	case 0, 1:
		temp = str
	case 2:
		temp = str[:1] + "*"
	case 3:
		temp = str[:1] + "*" + str[len(str)-1:]
	case 4:
		temp = str[:1] + "**" + str[len(str)-1:]
	case 5:
		temp = str[:2] + "**" + str[len(str)-1:]
	case 6:
		temp = str[:2] + "**" + str[len(str)-2:]
	case 7:
		temp = str[:2] + "***" + str[len(str)-2:]
	case 8:
		temp = str[:2] + "***" + str[len(str)-3:]
	case 9:
		temp = str[:3] + "***" + str[len(str)-3:]
	case 10:
		temp = str[:3] + "***" + str[len(str)-4:]
	default:
		temp = str[:4] + "***" + str[len(str)-4:]
	}
	return temp
}

// StrToCamel 驼峰命名法
func StrToCamel(str string) string {
	chunks := chunk(str)
	for i, c := range chunks {
		if i == 0 {
			chunks[i] = strings.ToLower(c)
			continue
		}
		chunks[i] = strings.Title(c)
	}
	return strings.Join(chunks, "")
}

// StrToKebab 短横线隔开式命名法(又名：spinal case 脊柱命名法)
func StrToKebab(str string) string {
	chunks := chunk(str)
	return strings.Join(chunks, "-")
}

// StrToPascal 帕斯卡命名法
func StrToPascal(str string) string {
	chunks := chunk(str)
	for i, c := range chunks {
		chunks[i] = strings.Title(c)
	}
	return strings.Join(chunks, "")
}

// StrToSnake 蛇形命名法
func StrToSnake(str string) string {
	chunks := chunk(str)
	return strings.Join(chunks, "_")
}

// StrReverse 反转字符串
func StrReverse(s string) string {
	runes := []rune(s)
	for from, to := 0, len(runes)-1; from < to; from, to = from+1, to-1 {
		runes[from], runes[to] = runes[to], runes[from]
	}
	return string(runes)
}
