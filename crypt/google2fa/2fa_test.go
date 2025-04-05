package google2fa

import (
	"testing"
)

func TestTotp(t *testing.T) {
	g := GoogleAuthenticator2FaSha1{
		Base32NoPaddingEncodedSecret: "NG5Y36KNEA2QCYVTUEPDMHIZ7YACRBWR",
		ExpireSecond:                 30,
		Digits:                       6,
	}

	totp, err := g.Totp()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(totp)
}

func TestQr(t *testing.T) {
	g := GoogleAuthenticator2FaSha1{
		Base32NoPaddingEncodedSecret: "NG5Y36KNEA2QCYVTUEPDMHIZ7YACRBWR",
		ExpireSecond:                 30,
		Digits:                       6,
	}
	qrString := g.QrString("X", "Mirage")
	t.Log(qrString)
}
