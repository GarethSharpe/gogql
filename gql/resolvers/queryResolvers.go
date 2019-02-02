package resolvers

import (
	"context"

	"github.com/garethsharpe/gql/utils"
	"github.com/garethsharpe/gql/models"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Case(ctx context.Context, id string) (models.Case, error) {
	accessToken := ctx.Value(utils.ACCESS_TOKEN).(string)
	returnCase, err := r.AppSvc.GetCase(accessToken, id)
	return returnCase, err
}

func (r *queryResolver) Cases(ctx context.Context) ([]models.Case, error) {
	accessToken := ctx.Value(utils.ACCESS_TOKEN).(string)
	cases, err := r.AppSvc.GetCases(accessToken)
	return cases, err
}