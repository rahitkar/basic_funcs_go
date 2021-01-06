package main

import (
	"fmt"
	"rsc.io/quote"
	"step.com/basic_funcs/src/closures"
	"step.com/basic_funcs/src/factorial"
	"step.com/basic_funcs/src/series"
	"step.com/basic_funcs/src/shape"
)

func main() {
	nextInt := closures.Incrementer()
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	square := closures.Squarer(2)
	fmt.Println(square())
	fmt.Println(square())

	fmt.Println("factorial of 5 = ", factorial.Calculate(5))

	result := series.Fibonacci(6)
	fmt.Println(result)

	rectangle := shape.Rectangle{Length: 4, Breath: 6}
	fmt.Println(shape.CalculateArea(rectangle))

	circle := shape.Circle{Radius: 2}
	fmt.Println(shape.CalculateArea(circle))

	fmt.Println(quote.Go())
}