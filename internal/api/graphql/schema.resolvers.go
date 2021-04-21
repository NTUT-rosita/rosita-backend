package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/NTUT-rosita/rosita-backend/internal/api/graphql/generated"
	"github.com/NTUT-rosita/rosita-backend/internal/api/graphql/model"
	"github.com/NTUT-rosita/rosita-backend/internal/pkg/sliceConverter"
)

func (r *queryResolver) DayStudentDetailsPossibleValue(ctx context.Context) (*model.DayStudentDetailsPossibleValue, error) {

	result := &model.DayStudentDetailsPossibleValue{}

	for _, field := range graphql.CollectAllFields(ctx) {
		switch field {
		case "collegeType":
			distinctCollegeType, err := r.Queries.GetDistinctCollegeType(ctx)
			if err != nil {
				return nil, err
			}
			result.CollegeType = sliceConverter.NullStringToString(distinctCollegeType)
		case "academicYear":
			distinctAcademicYear, err := r.Queries.GetDistinctAcademicYear(ctx)
			if err != nil {
				return nil, err
			}
			result.AcademicYear = sliceConverter.Int16ToInt(distinctAcademicYear)
		case "semester":
			distinctSemester, err := r.Queries.GetDistinctSemester(ctx)
			if err != nil {
				return nil, err
			}
			result.Semester = sliceConverter.Int16ToInt(distinctSemester)

		case "schoolSystem":
			distinctSchoolSystem, err := r.Queries.GetDistinctSchoolSystem(ctx)
			if err != nil {
				return nil, err
			}
			result.SchoolSystem = sliceConverter.NullStringToString(distinctSchoolSystem)
		case "departmentName":
			distinctDepartmentName, err := r.Queries.GetDistinctDepartmentName(ctx)
			if err != nil {
				return nil, err
			}
			result.DepartmentName = sliceConverter.NullStringToString(distinctDepartmentName)
		}
	}

	return result, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
