package aes

import "github.com/gzylg/kits/errs"

var (
	ErrKeyLengthSixteen = errs.New("a sixteen or twenty-four or thirty-two length secret key is required")
)
