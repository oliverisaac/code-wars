package kata

func Tribonacci(signature [3]float64, n int) []float64 {
	result := []float64{}
	result = append(result, signature[0], signature[1], signature[2])

	for len(result) < n {
		len := len(result)
		result = append(result, (result[len-3] + result[len-2] + result[len-1]))
	}

	return result[:n]
}
