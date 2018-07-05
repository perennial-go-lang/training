package main

type triangle struct {
	Base float64
	Height float64
}

func (t triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
