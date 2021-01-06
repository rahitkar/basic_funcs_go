package factorial

func Calculate(num int) int {
	mul := 1

	for i := 0; i < num; i++ {
		mul *= num - i
	}
	return mul
}