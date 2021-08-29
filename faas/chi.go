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

const defaultMainAddr = "0.0.0.0:8081"

func RunTencentChi(r *chi.Mux) {
	if deployutil.IsStageDev() {
		err := http.ListenAndServe(defaultMainAddr, r)
		if err != nil {
			panic(err)
		}

		return
	}

	ChiFaas = chiadapter.New(r)
	faas.Start(HandlerChi)
}

func RunTencentChiWithSwagger(r *chi.Mux) {
	fs := http.FileServer(http.Dir("docs"))
	r.Handle("/chi/docs/*", http.StripPrefix("/chi/docs/", fs))
	r.Handle("/docs/*", http.StripPrefix("/docs/", fs))

	if deployutil.IsStageDev() {
		err := http.ListenAndServe(defaultMainAddr, r)
		if err != nil {
			panic(err)
		}

		return
	}

	ChiFaas = chiadapter.New(r)
	faas.Start(HandlerChi)
}
