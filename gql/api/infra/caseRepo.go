package repos

import (
	"fmt"
	"math/rand"

	"gogql/gql/models"
	"gogql/gql/utils"
)

type CaseRepo struct {
	cases []models.Case
}

type ICaseRepo interface {
	GetCase(accessToken string, id string) (*models.Case, error)
	GetCases(accessToken string) (*[]models.Case, error)
	CreateCase(accessToken string, caseArg models.InputCase) (string, error)
}

type Response struct {
	Records []*models.Case
}

// GetCase gets a case by an id.
// @author	Gareth Sharpe
// @param	{string}			The access token of a community user.
// @returns	{Case[]}			An array of cases.
func (r *CaseRepo) GetCase(accessToken string, id string) (*models.Case, error) {
	var c models.Case
	connection, err := utils.GetConnection(accessToken)
	if err != nil { return nil, err }
	err = connection.GetSObject(id, nil, &c)
	if err != nil { return nil, err }
	return &c, nil
}

// GetCases gets a list of cases new case.
// @author	Gareth Sharpe
// @param	{string}			The access token of a community user.
// @returns	{Case[]}			An array of cases.
func (r *CaseRepo) GetCases(accessToken string) (*[]models.Case, error) {
	res := &Response{}
	var cases []models.Case
	var c models.Case
	connection, err := utils.GetConnection(accessToken)
	if err != nil {
		return &cases, err
	}
	fields := utils.GetFields(c)
	soql := fmt.Sprintf("select %s from %s", *fields, c.ApiName())
	err = connection.Query(soql, res)
	if err != nil {
		return &cases, err
	}
	nRecords := len(res.Records)
	cases = make([]models.Case, nRecords)
	for i := range cases {
		cases[i] = *res.Records[i]
	}
	return &cases, nil
}

// CreateCase creates a new case.
// @author	Gareth Sharpe
// @param	{string}			The access token of a community user.
// @param	{Case}				The case object to create.
// @returns	{string}			The id of the newly created case.
func (r *CaseRepo) CreateCase(accessToken string, caseArg models.InputCase) (string, error) {
	id := fmt.Sprintf("T%d", rand.Int())
	c := models.Case{
		Id:   id,
	}
	r.cases = append(r.cases, c)
	return id, nil
}
