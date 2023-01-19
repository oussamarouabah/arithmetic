package arithmetics

func Sum(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func Avg(nums ...int) float64 {
	return float64(Sum(nums...)) / float64(len(nums))
}
