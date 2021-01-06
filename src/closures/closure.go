package closures

//Incrementer
func Incrementer() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

//Squarer
func Squarer(base int) func() int {
	baseValue := base
	return func() int {
		baseValue *= baseValue
		return baseValue
	}
}
