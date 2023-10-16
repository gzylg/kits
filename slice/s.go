package slice

import "math/rand"

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

func ShuffleInt(arr []int) []int {
	nums := make([]int, len(arr))
	buf := make([]int, len(arr))
	copy(buf, arr)
	for i := range nums {
		j := rand.Intn(len(buf))
		nums[i] = buf[j]
		buf = append(buf[0:j], buf[j+1:]...)
	}
	return nums
}
