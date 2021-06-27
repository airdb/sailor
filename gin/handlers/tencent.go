package handlers

import (
	"context"

	"github.com/airdb/sailor/deployutil"
	"github.com/gin-gonic/gin"
	"github.com/serverless-plus/tencent-serverless-go/events"
	"github.com/serverless-plus/tencent-serverless-go/faas"

	ginAdapter "github.com/serverless-plus/tencent-serverless-go/gin"
)

var GinFaas *ginAdapter.GinFaas

// Handler serverless faas handler.
func Handler(ctx context.Context, req events.APIGatewayRequest) (events.APIGatewayResponse, error) {
	return GinFaas.ProxyWithContext(ctx, req)
}

func RunTencentServerless(r *gin.Engine) {
	if deployutil.GetDeployStage() == deployutil.DeployStageDev {
		defaultAddr := ":8081"
		err := r.Run(defaultAddr)
		if err != nil {
			panic(err)
		}

		return
	}

	GinFaas = ginAdapter.New(r)

	faas.Start(Handler)
}
