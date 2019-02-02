package resolvers

import (
	"context"

	"github.com/garethsharpe/gql/utils"
	"github.com/garethsharpe/gql/models"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateCase(ctx context.Context, caseArg models.InputCase) (string, error) {
	accessToken := ctx.Value(utils.ACCESS_TOKEN).(string)
	id, err := r.AppSvc.CreateCase(accessToken, caseArg)
	return id, err
}