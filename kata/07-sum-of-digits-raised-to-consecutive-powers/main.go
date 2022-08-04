package kata

import "math"

// https://www.codewars.com/kata/5626b561280a42ecc50000d1/train/go

// SumDigPow looks for magic numbers between a and b whose digis raised to the power of their position is equal to the number
func SumDigPow(start, end uint64) []uint64 {
	result := []uint64{}

	for i := start; i <= end; i++ {
		if i == doSum(i) {
			result = append(result, i)
		}
	}

	return result
}

func doSum(val uint64) (sum uint64) {
	var thisDigit uint64
	digits := []uint64{}
	for val > 0 {
		thisDigit = val % 10
		val = (val - thisDigit) / 10
		digits = append(digits, thisDigit)
	}

	var result uint64 = 0
	var power int = 1
	for i := len(digits) - 1; i >= 0; i-- {
		result = result + uint64(math.Pow(float64(digits[i]), float64(power)))
		power++
	}
	return result
}
