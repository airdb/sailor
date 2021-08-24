package faas

import (
	"context"
	"net/http"

	"github.com/airdb/sailor/deployutil"
	"github.com/go-chi/chi/v5"
	chiadapter "github.com/serverless-plus/tencent-serverless-go/chi"
	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"
)

var ChiFaas *chiadapter.ChiFaas

func HandlerChi(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	return ChiFaas.ProxyWithContext(ctx, req)
}

func RunTencentChi(r *chi.Mux) {
	if deployutil.GetDeployStage() == deployutil.DeployStageDev {
		defaultAddr := ":8081"
		err := http.ListenAndServe(defaultAddr, r)
		if err != nil {
			panic(err)
		}

		return
	}

	ChiFaas = chiadapter.New(r)
	faas.Start(HandlerChi)
}
