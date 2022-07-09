package param

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

/*
* GetParamArr 通过传入的数值得出对应的参数数组（如：3=1+2；7=1+2+4）
 */
func GetParamArr(p int) (arr []int) {
	sprintf := fmt.Sprintf("%b", p)
	for i, s := range sprintf {
		a := fmt.Sprintf("%c", s)
		if a == "1" {
			for j := 1; j < len(sprintf)-i; j++ {
				a += "0"
			}
			arr = append(arr, binDec(a))
		}

	}
	return
}

func binDec(b string) (n int) {
	s := strings.Split(b, "")
	l := len(s)
	i := 0
	d := float64(0)
	for i = 0; i < l; i++ {
		f, err := strconv.ParseFloat(s[i], 64)
		if err != nil {
			log.Println("Binary to decimal error:", err.Error())
			return -1
		}
		d += f * math.Pow(2, float64(l-i-1))
	}
	return int(d)
}
