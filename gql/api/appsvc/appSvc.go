package api

import (
	repos "gogql/gql/api/infra"
	"gogql/gql/models"
)

type AppSvc struct {
	CaseRepo repos.ICaseRepo
}

type IAppSvc interface {
	GetCase(accessToken string, id string) (*models.Case, error)
	GetCases(accessToken string) (*[]models.Case, error)
	CreateCase(accessToken string, caseArg models.InputCase) (string, error)
}

func GetAppSvcInstance() AppSvc {
	caseRepo := repos.CaseRepo{}
	appSvcInstance := AppSvc{
		CaseRepo: &caseRepo,
	}
	return appSvcInstance
}

func GetTestAppSvcInstance() AppSvc {
	mockCaseRepo := repos.MockCaseRepo{}
	appSvcInstance := AppSvc{
		CaseRepo: &mockCaseRepo,
	}
	return appSvcInstance
}

func (svc *AppSvc) GetCase(accessToken string, id string) (*models.Case, error) {
	returnCase, err := svc.CaseRepo.GetCase(accessToken, id)
	return returnCase, err
}

func (svc *AppSvc) GetCases(accessToken string) (*[]models.Case, error) {
	cases, err := svc.CaseRepo.GetCases(accessToken)
	return cases, err
}

func (svc *AppSvc) CreateCase(accessToken string, caseArg models.InputCase) (string, error) {
	id, err := svc.CaseRepo.CreateCase(accessToken, caseArg)
	return id, err
}
