package repos

import (
	"github.com/garethsharpe/gql/models"
	"github.com/garethsharpe/gql/utils"
)

type MockCaseRepo struct {
	cases []models.Case
}

func (r *MockCaseRepo) GetCase(accessToken string, id string) (*models.Case, error) {
	var c = models.Case{Id: utils.ID}
	return &c, nil
}

func (r *MockCaseRepo) GetCases(accessToken string) (*[]models.Case, error) {
	var c = models.Case{Id: utils.ID}
	var cases = []models.Case{c}
	return &cases, nil
}

func (r *MockCaseRepo) CreateCase(accessToken string, caseArg models.InputCase) (string, error) {
	id := utils.ID
	return id, nil
}
