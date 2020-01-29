package middlewares_test

import (
	"net/http"
	"net/http/httptest"
	"path"
	"testing"

	"github.com/airdb/sailor/enum"
	"github.com/airdb/sailor/gin/middlewares"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func APIClient(method, path string) (*httptest.ResponseRecorder, error) {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(
		middlewares.Jsonifier(),
	)

	router.GET("/success", func(c *gin.Context) {
		middlewares.SetResp(
			c,
			// enum.AirdbSuccess,
			enum.AirdbSuccess,
			"success",
		)
	})

	router.GET("/failed", func(c *gin.Context) {
		middlewares.SetResp(
			c,
			// enum.AirdbSuccess,
			enum.AirdbFailed,
			"failed",
		)
	})

	req, err := http.NewRequest(method, path, nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	return w, err
}

func TestResponse(t *testing.T) {
	asserts := assert.New(t)

	for _, testCase := range testCases {
		uri := path.Join(testCase.URI)
		t.Log("test case:", testCase.TestID, testCase.URI)
		resp, err := APIClient(testCase.Method, uri)
		t.Log(resp.Body.String())
		asserts.NoError(err)

		assert.Equal(t, testCase.ExpectedCode, resp.Code)
	}
}

//You could write the init logic like reset database code here
var testCases = []struct {
	TestID         int
	URI            string
	Method         string
	bodyData       string
	ExpectedCode   int
	responseRegexg string
	msg            string
}{
	//---------------------   Testing case register   ---------------------
	{
		1,
		"/success",
		"GET",
		`{"user":{"username": "xxx","email": "","password": ""}}`,
		http.StatusOK,
		"",
		"valid data and should return StatusCreated",
	},
	{
		2,
		"/failed",
		"GET",
		`{"user":{"username": "xxx","email": "","password": ""}}`,
		http.StatusOK,
		"",
		"valid data and should return StatusCreated",
	},
}
