package main

import (
	"fmt"
	"rsc.io/quote"
	"step.com/basic_funcs/src/factorial"
	"step.com/basic_funcs/src/series"
)

func main() {
	fmt.Println("factorial of 5 = ", factorial.Calculate(5))

	result := series.Fibonacci(6)
	fmt.Println(result)

	fmt.Println(quote.Go())
}