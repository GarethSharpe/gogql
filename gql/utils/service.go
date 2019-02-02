package utils

import (
	"github.com/garethsharpe/gql/api/appsvc"
	"github.com/garethsharpe/gql/api/infra"
)

func GetAppSvcInstance() api.IAppSvc {
	caseRepo := repos.CaseRepo{}
	appSvcInstance := api.AppSvc{
		CaseRepo: caseRepo,
	}
	return &appSvcInstance
}