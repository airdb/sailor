package middlewares

import (
	"fmt"
	"net/http"

	"github.com/airdb/sailor/enum"
	"github.com/gin-gonic/gin"
)

func SetResp(c *gin.Context, code uint, value interface{}) {
	c.Set(ContextCode, int(code))
	fmt.Println("aaaa0000", code)
	fmt.Println("=====")
	c.Set(ContextKeyResp, value)
}

func Jsonifier() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Process request.
		c.Next()

		resp := &Response{}

		shouldJsonify := false
		statusCode := http.StatusOK

		code := c.GetInt(ContextCode)
		fmt.Println("aaa", code)
		// Jsonify the response.
		value, exists := c.Get(ContextKeyResp)
		if exists {
			resp.Success = true
			resp.Code = uint(code)
			resp.Content = value
			resp.Message = enum.FormCode(uint(code))
			shouldJsonify = true
		} else {
			resp.Success = false
			resp.Code = uint(code)
			resp.Content = value
			resp.Error = enum.FormCode(uint(code))
			shouldJsonify = true
		}

		if shouldJsonify {
			c.JSON(statusCode, resp)
		}
	}
}
