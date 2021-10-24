package faas

import (
	"context"
	"net/http"
	"path/filepath"

	"github.com/airdb/sailor/deployutil"
	"github.com/airdb/sailor/version"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
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

func RunTencentChiWithSwagger(r *chi.Mux, project string) {
	path := filepath.Join("/", project)

	// Return the default root index.
	r.Get(path, HandleVersion)

	path = filepath.Join("/", project, "/")
	r.Get(path, HandleVersion)

	path = filepath.Join("/", project, "/", "docs", "/")

	fs := http.FileServer(http.Dir("docs"))
	r.Handle(path+"*", http.StripPrefix(path, fs))
	r.Handle("/docs/*", http.StripPrefix("/docs/", fs))

	RunTencentChi(r)
}

// VersionHandler - Returns version information
// @Summary Version handler.
// @Description Returns version information, like repo, build, runtime, env
// @Tags version
// @Accept  json
// @Produce  json
// @Success 200 {string} response "api response"
// @Router / [get]
func HandleVersion(w http.ResponseWriter, r *http.Request) {
	info := version.GetBuildInfo()
	info.Swagger = "docs/swagger.yaml"
	render.JSON(w, r, info)

	w.WriteHeader(http.StatusOK)
}
