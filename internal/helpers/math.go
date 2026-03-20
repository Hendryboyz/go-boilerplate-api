package helpers

import "math"

func RoundFloat(value float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(value*ratio) / ratio
}

func FloorFloat(value float64, precision int) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Floor(value*ratio) / ratio
}

func GetRatio(portion, total float64, precision int) float64 {
	if total == 0 {
		return 0
	}
	if precision > 0 {
		theRatio := RoundFloat(portion/total, precision)
		if theRatio == 1.0 && portion != total {
			return FloorFloat(portion/total, precision)
		}
		return theRatio
	}
	return portion / total
}
