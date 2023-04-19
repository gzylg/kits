package ip

import (
	"fmt"
	"testing"
)

func TestGetExternalIP1(t *testing.T) {
	ip, err := GetExternalIP1()
	if err != nil {
		t.Fail()
	}

	fmt.Println(ip)
}

func TestGetExternalIP2(t *testing.T) {
	ip, err := GetExternalIP2()
	if err != nil {
		t.Fail()
	}

	fmt.Println(ip)
}

func TestGetExternalIP3(t *testing.T) {
	ip, err := GetExternalIP3()
	if err != nil {
		t.Fail()
	}

	fmt.Println(ip)
}

func TestGetExternalIP4(t *testing.T) {
	ip, err := GetExternalIP4()
	if err != nil {
		t.Fail()
	}

	fmt.Println(ip)
}
