package arraysslices

func Sum(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func SumAll(nums ...[]int) []int {
	sums := make([]int, len(nums))
	for idx, n := range nums {
		sums[idx] = Sum(n)
	}
	return sums
}

func SumAllTails(nums ...[]int) []int {
	var result []int
	for _, n := range nums {
		if len(n) < 2 {
			result = append(result, 0)
		} else {
			tail := n[1:]
			result = append(result, Sum(tail))
		}
	}
	return result
}
