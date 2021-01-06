package series

// Fibonacci returns fiboSeries in from of an array upto the given limit
func Fibonacci(to int) []int {
	fiboSeries := make([]int, 0)

	pri := -1
	next := 1

	for i := 0; i < to; i++ {
		sum := pri + next
		fiboSeries = append(fiboSeries, sum)
		pri = next
		next = sum
	}
	return fiboSeries
}

// Fibonacci returns fiboSeries in from of an array upto the given limit
// func Fibonacci(series []int, to int) []int {
// 	length := len(series)
// 	if length >= to {
// 		return series
// 	}
// 	return Fibonacci(append(series, series[length - 1] + series[length - 2]), to);
// }