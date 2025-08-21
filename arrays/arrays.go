package arrays

func Sum(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}

func SumAll(slices ...[]int) []int {

	arr := make([]int, 0, len(slices))

	for _, slice := range slices {

		sum := 0

		for _, num := range slice {
			sum += num
		}
		arr = append(arr, sum)
	}
	return arr
}
