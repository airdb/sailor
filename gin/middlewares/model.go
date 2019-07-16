package middlewares

const (
	ContextKeyResp = "resp"
	ContextKeyErr  = "err"
)

type ErrorRep struct {
	Status *int   `json:"status"`
	Errmsg string `json:"errmsg"`
}

type Response struct {
	Version string      `json:"version"`
	Success bool        `json:"success"`
	Error   interface{} `json:"error,omitempty"`
	Result  interface{} `json:"result,omitempty"`
}
