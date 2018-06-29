// Package school implements a simple grade school
package school

import "sort"

// Grade is the structure for a grade
type Grade struct {
	grade    int
	students []string
}

// School is the type for a school
type School map[int]Grade

// New returns a new School
func New() *School {
	return &School{}
}

// Add adds a student to a grade
func (s School) Add(student string, g int) {
	grade, ok := s[g]

	if !ok {
		grade = Grade{g, make([]string, 0)}
	}

	grade.students = append(grade.students, student)
	sort.Strings(grade.students)
	s[g] = grade
}

// Grade returns grade g of the school
func (s School) Grade(grade int) []string {
	return s[grade].students
}

// Enrollment returns all grades of the school sorted by grade
func (s School) Enrollment() []Grade {
	var keys []int
	var grades []Grade

	for i := range s {
		keys = append(keys, i)
	}
	sort.Ints(keys)

	for _, i := range keys {
		grades = append(grades, s[i])
	}

	return grades
}
