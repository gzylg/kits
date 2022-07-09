package calc

import "math"

const (
	rad                          = math.Pi / 180.0
	earthRadiusMeter     float64 = 6371000 // 地球半径(单位：米)  6.371 * math.Pow(10, 6)
	earthRadiusKilometer float64 = 6371    // 地球半径(单位：千米) 6.371 * math.Pow(10, 3)
)

// DistanceMeter 给出两个经纬度，获取其距离 (单位：米)
func DistanceMeter(lat1, lng1, lat2, lng2 float64) float64 {
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * earthRadiusMeter
}

// DistanceKilometer 给出两个经纬度，获取其距离 (单位：千米)
func DistanceKilometer(lat1, lng1, lat2, lng2 float64) float64 {
	lat1 = lat1 * rad
	lng1 = lng1 * rad
	lat2 = lat2 * rad
	lng2 = lng2 * rad

	theta := lng2 - lng1
	dist := math.Acos(math.Sin(lat1)*math.Sin(lat2) + math.Cos(lat1)*math.Cos(lat2)*math.Cos(theta))

	return dist * earthRadiusKilometer
}
