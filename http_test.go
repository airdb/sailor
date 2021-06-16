package sailor_test

import (
	"fmt"
	"testing"

	"github.com/airdb/sailor"
)

type resp struct {
	Login string `json:"login"`
	URL   string `json:"url"`
}

func TestHTTPRequest(t *testing.T) {
	client := sailor.NewHTTPClient()

	client.SetDomain("api.github.com")
	client.SetEndpoint("/users/airdb")
	client.SetUserAgent("test/v0.0.1")
	// client.SetDebug()

	output := resp{}
	if err := client.HTTPRequest(client, &output); err != nil {
		fmt.Println(err)

		return
	}

	fmt.Println(&output)
}
