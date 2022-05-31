package des

import "github.com/gzylg/kits/errs"

var (
	ErrKeyLengtheEight = errs.New("a eight-length secret key is required")
	ErrIvDes           = errs.New("a eight-length ivdes key is required")
)
