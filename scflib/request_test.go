package scflib_test

import (
	"testing"

	"github.com/airdb/scf-mina-api/scflib"
	"github.com/tencentyun/scf-go-lib/events"
)

type Req struct {
	PageNo   uint `form:"pageNo"`
	PageSize uint `form:"pageSize"`
}

func TestBindQuery(t *testing.T) {
	qs := events.APIGatewayQueryString{"pageNo": []string{"1"}, "pageSize": []string{"10"}}

	var r Req

	err := scflib.BindQuery(qs, &r)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(r, r.PageNo)
}
