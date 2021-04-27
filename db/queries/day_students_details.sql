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

-- name: Count :one
SELECT Count(*) FROM day_students_details Where (college_type = @collegeType OR @collegeType is NULL) AND (academic_year = @academicYear OR NULLIF(@academicYear, 0) is NULL) AND(semester = @semester OR NULLIF(@semester,0) is NULL) AND(school_system = @schoolSystem OR @schoolSystem is NULL) AND(department_name = @departmentName OR @departmentName is NULL);