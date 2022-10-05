package slice

// HasString 检查传入的字符串是否存在于slice中
func HasString(v string, sl []string) bool {
	if len(sl) < 1 {
		return false
	}

	for _, vv := range sl {
		if vv == v {
			return true
		}
	}
	return false
}

func HasInt(i int, il []int) bool {
	if len(il) < 1 {
		return false
	}

	for _, ii := range il {
		if ii == i {
			return true
		}
	}
	return false
}

func HasUint32(i uint32, il []uint32) bool {
	if len(il) < 1 {
		return false
	}

	for _, ii := range il {
		if ii == i {
			return true
		}
	}
	return false
}
