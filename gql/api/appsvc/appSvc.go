package api

import (
	"github.com/garethsharpe/gql/models"
	"github.com/garethsharpe/gql/api/infra"
)

type AppSvc struct { 
	CaseRepo repos.CaseRepo
}

type IAppSvc interface {
	GetCase(accessToken string, id string) (models.Case, error)
	GetCases(accessToken string) ([]models.Case, error)
	CreateCase(accessToken string, caseArg models.InputCase) (string, error)
}

func (svc *AppSvc) GetCase(accessToken string, id string) (models.Case, error) {
	// domain logic
	returnCase, err := svc.CaseRepo.GetCase(accessToken, id)
	// domain logic
	return returnCase, err
}

func (svc *AppSvc) GetCases(accessToken string) ([]models.Case, error) {
	// domain logic
	cases, err := svc.CaseRepo.GetCases(accessToken)
	// domain logic
	return cases, err
}

func (svc *AppSvc) CreateCase(accessToken string, caseArg models.InputCase) (string, error) {
	//
	id, err := svc.CaseRepo.CreateCase(accessToken, caseArg)
	//
	return id, err
}
