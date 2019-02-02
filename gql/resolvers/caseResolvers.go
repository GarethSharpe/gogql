package resolvers

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/garethsharpe/gql/models"
)

type caseResolver struct{ *Resolver }


func (r *caseResolver) LastModifiedBy(ctx context.Context, c *models.Case) (*models.User, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	returnUser := models.User{ID: &id}
	return &returnUser, nil
}

func (r *caseResolver) Asset(ctx context.Context, c *models.Case) (*models.Asset, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	returnAsset := models.Asset{ID: &id}
	return &returnAsset, nil
}

func (r *caseResolver) Contact(ctx context.Context, c *models.Case) (*models.Contact, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	returnContact := models.Contact{ID: &id}
	return &returnContact, nil
}

func (r *caseResolver) CreatedBy(ctx context.Context, c *models.Case) (*models.User, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	returnUser := models.User{ID: &id}
	return &returnUser, nil
}

func (r *caseResolver) Owner(ctx context.Context, c *models.Case) (*models.User, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	returnUser := models.User{ID: &id}
	return &returnUser, nil
}