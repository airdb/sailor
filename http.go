package sailor

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

type HTTPClient struct {
	debug bool

	url    string
	method string
	// Implement `curl --data-urlencode`
	values  url.Values
	headers map[string]string

	body interface{}
}

func (client *HTTPClient) SetDebug() {
	client.debug = true
}

func (client *HTTPClient) GetDebug() bool {
	return client.debug
}

func (client *HTTPClient) SetURL(url string) {
	client.url = url
}

func (client *HTTPClient) GetURL() string {
	return client.url
}

func (client *HTTPClient) SetMethod(method string) {
	client.method = method
}

func (client *HTTPClient) GetMethod() string {
	return client.method
}

func (client *HTTPClient) SetHeaders(headers map[string]string) {
	client.headers = headers
}

func (client *HTTPClient) GetHeaders() map[string]string {
	return client.headers
}

func (client *HTTPClient) SetValues(values url.Values) {
	client.values = values
}

func (client *HTTPClient) GetValues() url.Values {
	return client.values
}

func (client *HTTPClient) SetBody(body interface{}) {
	client.body = body
}

func (client *HTTPClient) GetBody() interface{} {
	return client.body
}

type RequestInterface interface {
	SetDebug()
	GetDebug() bool

	SetURL(string)
	GetURL() string

	SetMethod(string)
	GetMethod() string

	SetHeaders(headers map[string]string)
	GetHeaders() map[string]string

	SetValues(url.Values)
	GetValues() url.Values

	SetBody(interface{})
	GetBody() interface{}
}

type ResponseInterface interface{}

var DefaultClient = &http.Client{
	Timeout: 60 * time.Second,
}

func HTTPRequest(requestInterface RequestInterface, responseInterface ResponseInterface) error {
	return HTTPRequestWithClient(DefaultClient, requestInterface, responseInterface)
}

func HTTPRequestWithClient(client *http.Client,
	requestInterface RequestInterface,
	responseInterface ResponseInterface,
) error {
	switch requestInterface.GetMethod() {
	case http.MethodGet:
		var r *strings.Reader

		if requestInterface.GetValues() != nil {
			r = strings.NewReader(requestInterface.GetValues().Encode())
		}

		req, err := http.NewRequest(requestInterface.GetMethod(), requestInterface.GetURL(), r)
		if err != nil {
			return err
		}

		if requestInterface.GetValues() != nil {
			req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		}

		// The RawQuery should binding on `url`.  Like follow:
		// Query string `url:"q"`
		body, err := query.Values(requestInterface.GetBody())
		if err != nil {
			return err
		}

		if len(body) > 0 {
			req.URL.RawQuery = body.Encode()
		}

		return DoRequest(client, req, requestInterface, responseInterface)
	case http.MethodPost, http.MethodPatch:
		payload, err := json.Marshal(requestInterface.GetBody())
		if err != nil {
			return err
		}

		req, err := http.NewRequest(requestInterface.GetMethod(), requestInterface.GetURL(), bytes.NewBuffer(payload))
		if err != nil {
			return err
		}

		req.Header.Add("Content-Type", "application/json")

		return DoRequest(DefaultClient, req, requestInterface, responseInterface)
	default:
	}

	return nil
}

func DoRequest(client *http.Client,
	req *http.Request,
	requestInterface RequestInterface,
	responseInterface ResponseInterface,
) error {
	for headerKey, headerValue := range requestInterface.GetHeaders() {
		req.Header.Set(headerKey, headerValue)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// For debug.
	if requestInterface.GetDebug() {
		log.Println(string(bodyBytes))
	}

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	switch resp.StatusCode {
	case http.StatusOK:
		return json.NewDecoder(resp.Body).Decode(&responseInterface)
	default:
		return errors.Errorf("http status code is not 200.")
	}
}
