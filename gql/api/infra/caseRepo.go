package repos

import (
	"fmt"
	"math/rand"
	"github.com/garethsharpe/gql/models"
)

type CaseRepo struct { 
	cases []models.Case
 }

type ICaseRepo interface {
	GetCase(accessToken string, id string) (models.Case, error)
	GetCases(accessToken string) ([]models.Case, error)
	CreateCase(accessToken string, caseArg models.InputCase)
}

func (r *CaseRepo) GetCase(accessToken string, id string) (models.Case, error) {
	var returnCase models.Case
	for _, c := range r.cases {
		if c.Id == id {
			returnCase = c
			break
		}
	}
	return returnCase, nil
}

func (r *CaseRepo) GetCases(accessToken string) ([]models.Case, error) {
	return r.cases, nil
}

func (r *CaseRepo) CreateCase(accessToken string, caseArg models.InputCase) (string, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	c := models.Case{
		Name:   *caseArg.Name,
		Id:     id,
	}
	r.cases = append(r.cases, c)
	return id, nil
}