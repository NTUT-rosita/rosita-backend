package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/NTUT-rosita/rosita-backend/internal/api/graphql/generated"
	"github.com/NTUT-rosita/rosita-backend/internal/api/graphql/model"
	"github.com/NTUT-rosita/rosita-backend/internal/pkg/sliceConverter"
)

func (r *queryResolver) DayStudentDetailsPossibleValue(ctx context.Context) (*model.DayStudentDetailsPossibleValue, error) {
	distinctAcademicYear, err := r.Queries.GetDistinctAcademicYear(ctx)

	if err != nil {
		return nil, err
	}

	return &model.DayStudentDetailsPossibleValue{AcademicYear: sliceConverter.Int16ToInt(distinctAcademicYear)}, err
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
