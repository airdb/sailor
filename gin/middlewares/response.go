package middlewares

import (
	"net/http"

	"github.com/airdb/sailor/enum"
	"github.com/gin-gonic/gin"
)

func SetResp(c *gin.Context, code uint, value interface{}) {
	c.Set(ContextCode, int(code))
	c.Set(ContextKeyResp, value)
}

func Jsonifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request.
		c.Next()

		statusCode := http.StatusOK

		code := uint(c.GetInt(ContextCode))
		if code == 0 {
			code = enum.AirdbUndefined
		}

		// Jsonify the response.
		value, exists := c.Get(ContextKeyResp)

		resp := &Response{
			Code:    code,
			Content: value,
		}
		if exists {
			resp.Success = true
			resp.Message = enum.FormCode(code)
		} else {
			resp.Success = false
			resp.Error = enum.FormCode(code)
		}

		c.JSON(statusCode, resp)
	}
}
