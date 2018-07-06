package students

type Student struct {
	Name       string
	Age        int
	Percentage float64
}

type Students []Student

type ByName Students

func (s ByName) Len() int {
	return len(s)
}

func (s ByName) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByName) Less(i, j int) bool {
	return s[i].Name < s[j].Name
}

type ByAge Students

func (s ByAge) Len() int {
	return len(s)
}

func (s ByAge) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByAge) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

type ByPercentage Students

func (s ByPercentage) Len() int {
	return len(s)
}

func (s ByPercentage) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByPercentage) Less(i, j int) bool {
	return s[i].Percentage < s[j].Percentage
}
