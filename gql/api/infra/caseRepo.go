package repos

import (
	"fmt"
	"math/rand"
	
	"github.com/garethsharpe/gql/models"
	"github.com/garethsharpe/gql/utils"
)

type CaseRepo struct { 
	cases []models.Case
 }

type ICaseRepo interface {
	GetCase(accessToken string, id string) (models.Case, error)
	GetCases(accessToken string) ([]models.Case, error)
	CreateCase(accessToken string, caseArg models.InputCase)
}

type Response struct {
	Records []*models.Case
}

func (r *CaseRepo) GetCase(accessToken string, id string) (models.Case, error) {
	var c models.Case
	connection, err := utils.GetConnection(accessToken);
	if err != nil { return c, err }
	err = connection.GetSObject(id, nil, c)
	if err != nil { return c, err }
	// for _, c := range r.cases {
	// 	if c.Id == id {
	// 		returnCase = c
	// 		break
	// 	}
	// }
	return c, nil
}

func (r *CaseRepo) GetCases(accessToken string) ([]models.Case, error) {
	res := &Response{}
	var cases []models.Case
	var c models.Case
	connection, err := utils.GetConnection(accessToken);
	if err != nil { return cases, err }
	fields := utils.GetFields(c)
	soql := fmt.Sprintf("select %s from %s", *fields, c.ApiName())
	err = connection.Query(soql, res)
	if err != nil { return cases, err }
	nRecords := len(res.Records)
	cases = make([]models.Case, nRecords)
	for i := range cases {
		cases[i] = *res.Records[i]
	}
	return cases, nil
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