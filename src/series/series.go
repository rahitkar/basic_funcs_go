package series
// Fibonacci returns fiboSeries in from of an array upto the given limit
// func Fibonacci(to int) []int {
// 	fiboSeries := make([]int, 0)
// 	fiboSeries = append(fiboSeries, 0)
// 	fiboSeries = append(fiboSeries, 1)

// 	for i := 2; i < to; i++ {
// 		fiboSeries = append(fiboSeries, (fiboSeries[i - 1] + fiboSeries[i - 2])) 
// 	}
// 	return fiboSeries;
// }

// Fibonacci returns fiboSeries in from of an array upto the given limit
func Fibonacci(series []int, to int) []int {
	length := len(series)
	if length >= to {
		return series
	}
	return Fibonacci(append(series, series[length - 1] + series[length - 2]), to);
}