package calc

import "math"

// Round 四舍五入，传入需要处理的小数，以及需要保留的位数
func Round(f float64, n int) float64 {
	pow10_n := math.Pow10(n)
	return math.Trunc((f+0.5/pow10_n)*pow10_n) / pow10_n
}
