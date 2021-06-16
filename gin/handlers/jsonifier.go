package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	RetCode uint        `json:"ret_code"`
	Message interface{} `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func Jsonifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request.
		c.Next()

		resp := &Response{}

		shouldJsonify := false

		// Jsonify the response.
		value, exists := c.Get(ResponseKey)
		if exists {
			resp.Success = true
			resp.Data = value
			resp.Message = nil
			shouldJsonify = true
		}

		value, exists = c.Get(ResponseKeyErr)
		if exists {
			// err := value

			resp.Success = false
			resp.Data = nil
			resp.Message = value
			shouldJsonify = true
		}

		// Swagger and other html request should not be jsonified.
		if shouldJsonify {
			c.JSON(http.StatusOK, resp)
		}
	}
}

const (
	ResponseKey    = "resp"
	ResponseKeyErr = "respErr"
)

func SetResp(c *gin.Context, value interface{}) {
	c.Set(ResponseKey, value)
}

func SetRespErr(c *gin.Context, err interface{}) {
	c.Set(ResponseKeyErr, err)
}
