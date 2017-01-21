package util

import "math"

// Round rounds a float val, using the roundOn val(0-1) to determine when to round, places determines how many decimal places
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// RoundToInt rounds float to int using 0.5 as the round on val
func RoundToInt(val float64) (newVal int) {
	return int(Round(val, 0.5, 0))
}
