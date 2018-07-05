package student

func (s studentSorter) Len() int {
	return len(s.students)
}

func (s studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func (s studentSorter) Less(i, j int) bool {
	return s.by(s.students[i], s.students[j])
}