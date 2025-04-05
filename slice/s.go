package slice

import (
	"math/rand"
	"time"
)

// FindIndex 在切片中查找指定值的索引，如果找到则返回索引，否则返回 -1
func FindIndex[T comparable](slice []T, val T) int {
	for idx, v := range slice {
		if v == val {
			return idx
		}
	}
	return -1
}

// Shuffle 随机打乱切片的排列顺序，函数支持任何类型的切片
func Shuffle[T any](arr []T) []T {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 初始化随机数生成器
	n := len(arr)
	for i := n - 1; i > 0; i-- {
		j := r.Intn(i + 1)              // 生成一个 0 到 i 的随机索引
		arr[i], arr[j] = arr[j], arr[i] // 交换元素
	}
	return arr
}

// Unique 去重并返回新切片，使用comparable约束，支持所有可用于比较的类型。struct等自定义类型，需确保其字段均为可比较类型
func Unique[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(input))
	for _, v := range input {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// UniqueInOrder 去重+排序并返回新切片，使用comparable约束，支持所有可用于比较的类型。struct等自定义类型，需确保其字段均为可比较类型
func UniqueInOrder[T comparable](input []T) []T {
	seen := make(map[T]struct{})
	result := make([]T, 0, len(input))
	for _, v := range input {
		if _, ok := seen[v]; !ok {
			seen[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}
