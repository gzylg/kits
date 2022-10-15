package random

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// UUIDv4 uuid
func UUIDv4() string {
	return uuid.NewV4().String()
}

// UUIDv4WithUpper 大写的uuid
func UUIDv4WithUpper() string {
	return strings.ToUpper(UUIDv4())
}

// UUIDv4WithUpperAndJustLetter 大写的uuid并去除中间的连接符"-"
func UUIDv4WithUpperAndJustLetter() string {
	s := UUIDv4WithUpper()
	return strings.ReplaceAll(s, "-", "")
}
