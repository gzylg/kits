package random

import (
	"log"
	"testing"
)

func TestRandomCreateBytes(t *testing.T) {
	for i := 0; i < 50; i++ {
		s := randomCreateBytes(32)
		log.Println(string(s))
	}
}
