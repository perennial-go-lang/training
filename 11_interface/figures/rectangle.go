package figures

type Rectanagle struct {
	Length float64
	Width float64
}

func (r Rectanagle) Area() float64 {
	return r.Length * r.Width
}
