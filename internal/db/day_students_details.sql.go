// Code generated by sqlc. DO NOT EDIT.
// source: day_students_details.sql

package db

import (
	"context"
	"database/sql"
)

const count = `-- name: Count :one
SELECT Count(*) FROM day_students_details Where (college_type = $1 OR $1 is NULL) AND (academic_year = $2 OR NULLIF($2, 0) is NULL) AND(semester = $3 OR NULLIF($3,0) is NULL) AND(school_system = $4 OR $4 is NULL) AND(department_name = $5 OR $5 is NULL)
`

type CountParams struct {
	Collegetype    sql.NullString `json:"collegetype"`
	Academicyear   int16          `json:"academicyear"`
	Semester       int16          `json:"semester"`
	Schoolsystem   sql.NullString `json:"schoolsystem"`
	Departmentname sql.NullString `json:"departmentname"`
}

func (q *Queries) Count(ctx context.Context, arg CountParams) (int64, error) {
	row := q.queryRow(ctx, q.countStmt, count,
		arg.Collegetype,
		arg.Academicyear,
		arg.Semester,
		arg.Schoolsystem,
		arg.Departmentname,
	)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const getDistinctAcademicYear = `-- name: GetDistinctAcademicYear :many
SELECT DISTINCT academic_year FROM day_students_details
`

func (q *Queries) GetDistinctAcademicYear(ctx context.Context) ([]int16, error) {
	rows, err := q.query(ctx, q.getDistinctAcademicYearStmt, getDistinctAcademicYear)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int16
	for rows.Next() {
		var academic_year int16
		if err := rows.Scan(&academic_year); err != nil {
			return nil, err
		}
		items = append(items, academic_year)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDistinctCollegeType = `-- name: GetDistinctCollegeType :many
SELECT DISTINCT college_type FROM day_students_details
`

func (q *Queries) GetDistinctCollegeType(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.query(ctx, q.getDistinctCollegeTypeStmt, getDistinctCollegeType)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var college_type sql.NullString
		if err := rows.Scan(&college_type); err != nil {
			return nil, err
		}
		items = append(items, college_type)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDistinctDepartmentName = `-- name: GetDistinctDepartmentName :many
SELECT DISTINCT department_name FROM day_students_details
`

func (q *Queries) GetDistinctDepartmentName(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.query(ctx, q.getDistinctDepartmentNameStmt, getDistinctDepartmentName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var department_name sql.NullString
		if err := rows.Scan(&department_name); err != nil {
			return nil, err
		}
		items = append(items, department_name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDistinctSchoolSystem = `-- name: GetDistinctSchoolSystem :many
SELECT DISTINCT school_system FROM day_students_details
`

func (q *Queries) GetDistinctSchoolSystem(ctx context.Context) ([]sql.NullString, error) {
	rows, err := q.query(ctx, q.getDistinctSchoolSystemStmt, getDistinctSchoolSystem)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []sql.NullString
	for rows.Next() {
		var school_system sql.NullString
		if err := rows.Scan(&school_system); err != nil {
			return nil, err
		}
		items = append(items, school_system)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getDistinctSemester = `-- name: GetDistinctSemester :many
SELECT DISTINCT semester FROM day_students_details
`

func (q *Queries) GetDistinctSemester(ctx context.Context) ([]int16, error) {
	rows, err := q.query(ctx, q.getDistinctSemesterStmt, getDistinctSemester)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []int16
	for rows.Next() {
		var semester int16
		if err := rows.Scan(&semester); err != nil {
			return nil, err
		}
		items = append(items, semester)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
