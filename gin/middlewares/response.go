package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetResp(c *gin.Context, value interface{}) {
	c.Set(ContextKeyResp, value)
}

func ToJSON(version string) gin.HandleFunc {
	return func(c *gin.Context) {
		// Process request.
		c.Next()

		resp := &Response{
			Version: version,
		}

		shouldJsonify := false
		statusCode := http.StatusOK

		// Jsonify the response.
		value, exists := c.Get(ContextKeyResp)
		if exists {
			resp.Success = true
			resp.Result = value
			resp.Error = nil
			shouldJsonify = true
		}

		if shouldJsonify {
			c.JSON(statusCode, resp)
		}
	}
}
