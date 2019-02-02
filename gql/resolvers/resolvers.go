package resolvers

import (
	"github.com/garethsharpe/gql/api/appsvc"
	"github.com/garethsharpe/gql/generated"
)



type Resolver struct{ AppSvc api.IAppSvc }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Case() generated.CaseResolver {
	return &caseResolver{r}
}


