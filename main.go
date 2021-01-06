package main

import (
	"fmt"
	"./src/factorial"
	"./src/series"
)

func main() {
	fmt.Println("factorial of 5 = ", factorial.Calculate(5))

	fiboSeries := make([]int, 0)
	fiboSeries = append(fiboSeries, 0)
	fiboSeries = append(fiboSeries, 1)
	result := series.Fibonacci(fiboSeries, 6)
	fmt.Println(fiboSeries)
	fmt.Println(result)
}