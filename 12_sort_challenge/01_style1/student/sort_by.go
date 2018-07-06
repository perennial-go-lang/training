package student

import "sort"

type By func(s1, s2 Student) bool

func (b By) SortAsc(students Students) {
	s := studentSorter{
		students,
		b,
	}
	sort.Sort(s)
}

func (b By) SortDesc(students Students) {
	s := studentSorter{
		students,
		b,
	}
	sort.Sort(sort.Reverse(s))
}

var Name = func(s1, s2 Student) bool {
	return s1.Name < s2.Name
}

var Age = func(s1, s2 Student) bool {
	return s1.Age < s2.Age
}

var Percentage = func(s1, s2 Student) bool {
	return s1.Percentage < s2.Percentage
}