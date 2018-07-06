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

func (s StudentParser) Less() int {
	return s.students.
}

func (by By) SortAsc(students Students) {
	sp := StudentParser{
		students,
		by,
	}
	sort.Sort(sp)
}
