package param

func GetArrUint32(p uint32) (arr []uint32) {
	for i := 0; p > 0; p >>= 1 {
		if p&1 == 1 {
			arr = append(arr, 1<<i)
		}
		i++
	}
	return
}
func GetArrInt(p uint32) (arr []uint32) {
	for i := 0; p > 0; p >>= 1 {
		if p&1 == 1 {
			arr = append(arr, 1<<i)
		}
		i++
	}
	return
}
