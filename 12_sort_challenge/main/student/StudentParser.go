package student

import (
	"sort"
)

type By func(s1, s2 Student) bool

type StudentParser struct {
	students Students
	by       By
}

func (s StudentParser) Len() int {
	return len(s.students)
}

func (s StudentParser) Less(i, j int) bool {
	return s.by(s.students[i], s.students[j])
}

func (s StudentParser) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func (by By) SortAsc(students Students) {
	sp := StudentParser{
		students,
		by,
	}
	sort.Sort(sp)
}

func (by By) SortDesc(students Students) {
	sp := StudentParser{
		students,
		by,
	}
	sort.Sort(sort.Reverse(sp))
}
