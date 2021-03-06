// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type DayStudentDetailsInput struct {
	CollegeType    *string `json:"collegeType"`
	AcademicYear   *int    `json:"academicYear"`
	Semester       *int    `json:"semester"`
	SchoolSystem   *string `json:"schoolSystem"`
	DepartmentName *string `json:"departmentName"`
}

type DayStudentDetailsPossibleValue struct {
	CollegeType    []string `json:"collegeType"`
	AcademicYear   []int    `json:"academicYear"`
	Semester       []int    `json:"semester"`
	SchoolSystem   []string `json:"schoolSystem"`
	DepartmentName []string `json:"departmentName"`
}
