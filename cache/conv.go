package cache

import (
	"fmt"

	"github.com/gzylg/kits/errs"
)

func GetBytes(reply any) ([]byte, error) {

	switch result := reply.(type) {
	case string:
		return []byte(result), nil
	case []byte:
		return result, nil
	case nil:
		return []byte{}, nil
	}

	return nil, errs.New(fmt.Sprintf("redis: unexpected type for Bytes, got type %T", reply))
}
