package arraysslices

func Sum(numbers []int) int {
	sum_value := 0
	for _, value := range numbers {
		sum_value += value
	}
	return sum_value
}

func SumAll(slices ...[]int) []int {
	var sums []int
	for _, elem := range slices {
		sums = append(sums, Sum(elem))
	}
	return sums
}
