package appSvc_test

import (
	"reflect"
	"testing"

	api "github.com/garethsharpe/gql/api/appsvc"
	repos "github.com/garethsharpe/gql/api/infra"
	"github.com/garethsharpe/gql/models"
	"github.com/garethsharpe/gql/utils"
)

func TestAppSvc_GetCase(t *testing.T) {
	type fields struct {
		CaseRepo repos.ICaseRepo
	}
	type args struct {
		accessToken string
		id          string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *models.Case
		wantErr bool
	}{
		{"GetCase()", fields{&repos.MockCaseRepo{}}, args{utils.ACCESS_TOKEN, utils.ID}, &models.Case{Id: utils.ID}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			svc := &api.AppSvc{
				CaseRepo: tt.fields.CaseRepo,
			}
			got, err := svc.GetCase(tt.args.accessToken, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppSvc.GetCase() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AppSvc.GetCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
