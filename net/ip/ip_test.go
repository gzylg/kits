package ip

import (
	"testing"
)

func TestIp2LongAndLong2ip(t *testing.T) {
	testIp := "192.168.0.1"

	uintIP := IPStrToUInt32(testIp)
	t.Log("IPStrToUInt32: ", uintIP)

	ip := UInt32ToIPStr(uintIP)
	t.Log("UInt32ToIPStr: ", ip)

	if testIp == ip {
		t.Log("test success")
	} else {
		t.Error("test failed")
	}
}

func TestAddrToUint32(t *testing.T) {
	// addr := AddrToUint32()
}
