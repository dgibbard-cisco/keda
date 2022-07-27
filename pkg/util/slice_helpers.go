package util

// Takes a slice of strings, and looks for a string in it. If found it will
// return it's key/index, otherwise it will return -1 and a bool of false.
func FindStringInSlice(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

// Find the largest value in a slice of floats
func MaxFloatFromSlice(results []float64) float64 {
	max := results[0]
	for _, result := range results {
		if result > max {
			max = result
		}
	}
	return max
}

// Find the average value in a slice of floats
func AvgFloatFromSlice(results []float64) float64 {
	total := 0.0
	for _, result := range results {
		total += result
	}
	return total / float64(len(results))
}
