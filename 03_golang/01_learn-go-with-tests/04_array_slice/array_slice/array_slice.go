package array_slice

// SumList calculates sum of items in an array
func SumList(numbers [5]int) int {
	result := 0
	for _, v := range numbers {
		result += v
	}
	return result
}

