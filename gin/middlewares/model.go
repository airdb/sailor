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

type Error struct {
	StatusCode *int      `json:"-"`
	Message    string    `json:"message"`
	Type       ErrorType `json:"type"`
	Traceback  string    `json:"traceback"`
}

type ErrorType string
