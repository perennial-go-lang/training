package student

import "sort"

type Student struct {
	Name string
	Age int
	Percentage float64
}

type Students []Student

///////////

type by func(s1, s2 *Student) bool

type studentSorter struct {
	students Students
	by []by
}

func (s *studentSorter) Sort(students Students) {
	s.students = students
	sort.Sort(s)
}

func By(by ...by) *studentSorter {
	return &studentSorter{
		by: by,
	}
}

func (s *studentSorter) Len() int {
	return len(s.students)
}

func (s *studentSorter) Swap(i, j int) {
	s.students[i], s.students[j] = s.students[j], s.students[i]
}

func (s *studentSorter) Less(i, j int) bool {
	s1, s2 := &s.students[i], &s.students[j]

	var k int
	for k = 0; k < len(s.by) - 1; k++ {
		by := s.by[k]
		switch {
		case by(s1, s2):
			return true
		case by(s2, s1):
			return false
		}
	}
	return s.by[k](s1, s2)
}

var Name = func(s1, s2 *Student) bool {
	return s1.Name < s2.Name
}

var Age = func(s1, s2 *Student) bool {
	return s1.Age < s2.Age
}

var Percentage = func(s1, s2 *Student) bool {
	return s1.Percentage < s2.Percentage
}