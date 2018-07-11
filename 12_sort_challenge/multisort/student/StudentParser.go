package student

import (
	"sort"
)

type singleBy func(s1, s2 Student) bool

type By []singleBy

type StudentParser struct {
	students Students
	by       By
}

func (s StudentParser) Len() int {
	return len(s.students)
}

func (s StudentParser) Less(i, j int) (status bool) {
	for p, _ := range s.by {
		status := s.by[p](s.students[i], s.students[j])
		if status != s.by[p](s.students[j], s.students[i]) {
			return status
		}
	}
	return status
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
