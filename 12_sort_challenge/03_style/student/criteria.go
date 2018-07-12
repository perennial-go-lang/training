package student

type ByAge Students

func (b ByAge) Len() int {
	return len(b)
}

func (b ByAge) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByAge) Less(i, j int) bool {
	return b[i].Age < b[j].Age
}

type ByPercentage Students

func (b ByPercentage) Len() int {
	return len(b)
}

func (b ByPercentage) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByPercentage) Less(i, j int) bool {
	return b[i].Percentage < b[j].Percentage
}
