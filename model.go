package sailor

type HTTPAirdbResponse struct {
	Code    uint        `json:"code"`
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error,omitempty"`
	Message interface{} `json:"message,omitempty"`
}
