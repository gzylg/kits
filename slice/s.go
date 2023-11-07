package slice

import "math/rand"

func FindIndex[T comparable](slice []T, val T) int {
	for idx, v := range slice {
		if v == val {
			return idx
		}
	}
	return -1
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

func ShuffleUint32(arr []uint32) []uint32 {
	nums := make([]uint32, len(arr))
	buf := make([]uint32, len(arr))
	copy(buf, arr)
	for i := range nums {
		j := rand.Intn(len(buf))
		nums[i] = buf[j]
		buf = append(buf[0:j], buf[j+1:]...)
	}
	return nums
}
