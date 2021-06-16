package acs_test

import (
	"testing"
	"time"

	"github.com/airdb/sailor/acs"
)

type T struct {
	DeployInfo struct {
		GoVersion string    `json:"GoVersion"`
		Env       string    `json:"Env"`
		Repo      string    `json:"Repo"`
		Version   string    `json:"Version"`
		Build     string    `json:"Build"`
		BuildTime string    `json:"BuildTime"`
		CreatedAt time.Time `json:"CreatedAt"`
	} `json:"deploy_info"`
}

func TestNewClient(t *testing.T) {
	client, _ := acs.NewClientWithoutConfig()

	req := acs.BaseRequest{}
	req.SetDomain("scf.baobeihuijia.com")
	req.SetEndpoint("/test/wechat/")
	req.SetDebug()

	// var resp acs.Response
	var resp T
	err := client.DoAction(&req, &resp)
	t.Log(err, resp)
}
