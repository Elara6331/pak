package main

// Check if slice contains string
func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func Max(array []float64) float64 {
	var max = array[0]
	var min = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return max
}

// Get key from map given value
func GetKey(inMap map[string]string, val string) string {
	// For every key/value pair in map
	for key, value := range inMap {
		// If value found
		if value == val {
			// Return key
			return key
		}
	}
	// If fails, return empty string
	return ""
}

// Get slice of float64 given map[string]float64
func GetValuesDist(inMap map[string]float64) []float64 {
	// Make new slice with set capacity
	values := make([]float64, len(inMap))
	// Set index to 0
	index := 0
	for _, value := range inMap {
		// Set index of slice to value
		values[index] = value
		// Increment index
		index++
	}
	// Return completed slice
	return values
}
