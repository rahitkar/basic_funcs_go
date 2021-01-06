package shape

type Shape interface {
	Area() float64
}

func CalculateArea(s Shape) float64 {
	return s.Area()
}