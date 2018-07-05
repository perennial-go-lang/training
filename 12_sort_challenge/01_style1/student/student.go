package student

type Student struct {
	Name string
	Age int
	Percentage float64
}

type Students []Student

func (s Students) Len() int { return len(s)}
func (s Students) Swap(i, j int ) { s[i], s[j] = s[j], s[i] }

type ByName struct {
	Students
	Asc bool
}

func (b ByName) Less(i, j int) bool {
	if b.Asc {
		return b.Students[i].Name < b.Students[j].Name
	} else {
		return b.Students[i].Name > b.Students[j].Name
	}
}