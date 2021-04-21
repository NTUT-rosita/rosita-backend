-- name: GetDistinctCollegeType :many
SELECT DISTINCT college_type FROM day_students_details;

-- name: GetDistinctAcademicYear :many
SELECT DISTINCT academic_year FROM day_students_details;

-- name: GetDistinctSemester :many
SELECT DISTINCT semester FROM day_students_details;

-- name: GetDistinctSchoolSystem :many
SELECT DISTINCT school_system FROM day_students_details;

-- name: GetDistinctDepartmentName :many
SELECT DISTINCT department_name FROM day_students_details;