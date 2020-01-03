package middlewares

import (
	"net/http"

	"github.com/airdb/sailor/enum"
	"github.com/gin-gonic/gin"
)

func SetResp(c *gin.Context, status enum.Code, value interface{}) {
	// Set value must be int.
	c.Set(StatusCode, int(status))
	c.Set(ContextKeyResp, value)
}

func Jsonifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request.
		c.Next()

		resp := &Response{}

		shouldJsonify := false
		statusCode := http.StatusOK

		status := c.GetInt(StatusCode)
		// Jsonify the response.
		value, exists := c.Get(ContextKeyResp)
		if exists {
			resp.Success = true
			resp.Code = uint(status)
			resp.Data = value
			resp.Message = enum.FormCode(enum.Code(status))
			shouldJsonify = true
		} else {
			resp.Success = false
			resp.Code = uint(status)
			resp.Data = value
			resp.Error = enum.FormCode(enum.Code(status))
			shouldJsonify = true
		}

		if shouldJsonify {
			c.JSON(statusCode, resp)
		}
	}
}
