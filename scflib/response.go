package scflib

import (
	"github.com/tencentyun/scf-go-lib/events"
)

func ToHTML(statusCode int, msg string) (resp events.APIGatewayResponse) {
	resp.StatusCode = statusCode
	resp.Headers = map[string]string{"Content-Type": "text/html"}
	resp.Body = msg

	return resp
}
