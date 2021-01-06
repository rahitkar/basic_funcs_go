package shape

type Rectangle struct {
	Length float64
	Breath float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Breath
}