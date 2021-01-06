package main

import (
	"fmt"
	"rsc.io/quote"
	"step.com/main/src/factorial"
	"step.com/main/src/series"
)

func main() {
	fmt.Println("factorial of 5 = ", factorial.Calculate(5))

	result := series.Fibonacci(6)
	fmt.Println(result)

	fmt.Println(quote.Go())
}