package main

// Remove an element at an index from a slice
func removeAtIndex(s []string, index int) []string {
	return append(s[:index], s[index+1:]...)
}

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

func Find(slice []string, val string) int {
	for i, item := range slice {
		if item == val {
			return i
		}
	}
	return -1
}