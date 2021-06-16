package sailor

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
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
	debug        bool
	isSchemeHTTP bool
	timeout      int

	domain   string
	endpoint string
	url      string
	method   string

	// Implement `curl --data-urlencode`
	values  url.Values
	headers map[string]string

	body interface{}
}

const (
	UserAgent    = "sailor-agent/1.0.0"
	UserAgentKey = "User-Agent"
	SchemeHTTP   = "http://"
	SchemeHTTPS  = "https://"
)

func (client *HTTPClient) SetDebug() {
	client.debug = true
}

func (client *HTTPClient) GetDebug() bool {
	return client.debug
}

func (client *HTTPClient) SetDomain(domain string) {
	// Trim scheme.
	domain = strings.TrimPrefix(domain, SchemeHTTPS)
	domain = strings.TrimPrefix(domain, SchemeHTTP)

	client.domain = domain
}

func (client *HTTPClient) GetDomain() string {
	return client.domain
}

func (client *HTTPClient) SetEndpoint(endpoint string) {
	client.endpoint = endpoint
}

func (client *HTTPClient) GetEndpoint() string {
	return client.endpoint
}

func (client *HTTPClient) SetURL(url string) {
	client.url = url
}

func (client *HTTPClient) GetURL() string {
	if client.url != "" {
		return client.url
	}

	if client.isSchemeHTTP {
		client.url = SchemeHTTP + client.GetDomain() + client.GetEndpoint()
	} else {
		client.url = SchemeHTTPS + client.GetDomain() + client.GetEndpoint()
	}

	return client.url
}

func (client *HTTPClient) SetTimeout(timeout int) {
	client.timeout = timeout
}

func (client *HTTPClient) GetTimeout() int {
	return client.timeout
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

func (client *HTTPClient) SetUserAgent(ua string) {
	if client.headers == nil {
		header := make(map[string]string)
		client.headers = header
	}

	client.headers[UserAgentKey] = ua
}

func (client *HTTPClient) GetUserAgent() string {
	return client.headers[UserAgentKey]
}

type RequestInterface interface {
	SetDebug()
	GetDebug() bool

	SetDomain(string)
	GetDomain() string

	SetEndpoint(string)
	GetEndpoint() string

	SetURL(string)
	GetURL() string

	SetMethod(string)
	GetMethod() string

	SetTimeout(int)
	GetTimeout() int

	SetHeaders(headers map[string]string)
	GetHeaders() map[string]string

	SetValues(url.Values)
	GetValues() url.Values

	SetBody(interface{})
	GetBody() interface{}

	SetUserAgent(string)
	GetUserAgent() string
}

type ResponseInterface interface{}

const defaultHTTPTimout = 60

func NewHTTPClient() *HTTPClient {
	return &HTTPClient{
		timeout: defaultHTTPTimout,
	}
}

func (client *HTTPClient) HTTPRequest(requestInterface RequestInterface, responseInterface ResponseInterface) error {
	if requestInterface.GetURL() == "" {
		return errors.New("API url is null, please check")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(defaultHTTPTimout))
	defer cancel()

	switch requestInterface.GetMethod() {
	case http.MethodGet:
		return client.HTTPRequestGet(ctx, requestInterface, responseInterface)
	case http.MethodPost, http.MethodPatch:
		return client.HTTPRequestPost(ctx, requestInterface, responseInterface)
	default:
		requestInterface.SetMethod(http.MethodGet)

		return client.HTTPRequestGet(ctx, requestInterface, responseInterface)
	}
}

func (client *HTTPClient) HTTPRequestGet(
	ctx context.Context,
	requestInterface RequestInterface,
	responseInterface ResponseInterface,
) error {
	/*
		var r *strings.Reader

		if requestInterface.GetValues() != nil {
			r = strings.NewReader(requestInterface.GetValues().Encode())
		}
		log.Println(r)
	*/
	request, err := http.NewRequestWithContext(ctx,
		requestInterface.GetMethod(),
		requestInterface.GetURL(),
		nil)
	if err != nil {
		client.debugLog(err)

		return err
	}

	if requestInterface.GetValues() != nil {
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	}

	// The RawQuery should binding on `url`.  Like follow:
	// Query string `url:"q"`
	body, err := query.Values(requestInterface.GetBody())
	if err != nil {
		return err
	}

	if len(body) > 0 {
		request.URL.RawQuery = body.Encode()
	}

	return client.DoRequest(request, requestInterface, responseInterface)
}

func (client *HTTPClient) HTTPRequestPost(
	ctx context.Context,
	requestInterface RequestInterface,
	responseInterface ResponseInterface,
) error {
	payload, err := json.Marshal(requestInterface.GetBody())
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(
		ctx,
		requestInterface.GetMethod(),
		requestInterface.GetURL(),
		bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	request.Header.Add("Content-Type", "application/json")

	return client.DoRequest(request, requestInterface, responseInterface)
}

func (client *HTTPClient) DoRequest(
	request *http.Request,
	requestInterface RequestInterface,
	responseInterface ResponseInterface,
) error {
	for headerKey, headerValue := range requestInterface.GetHeaders() {
		request.Header.Set(headerKey, headerValue)
	}

	headers := requestInterface.GetHeaders()
	if headers[UserAgentKey] == "" {
		request.Header.Set(UserAgentKey, UserAgent)
	}

	_client := http.DefaultClient

	resp, err := _client.Do(request)
	if err != nil {
		client.debugLog(err)

		return err
	}
	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// For debug.
	client.debugLog(string(bodyBytes))

	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	switch resp.StatusCode {
	case http.StatusOK:
		return json.NewDecoder(resp.Body).Decode(&responseInterface)
	default:
		client.debugLog(resp.StatusCode)

		return errors.Errorf("http status code is not 200.")
	}
}

func (client *HTTPClient) debugLog(msg interface{}) {
	if client.debug {
		log.Println(msg)
	}
}

func (client *HTTPClient) Downloader() (io.ReadCloser, error) {
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, client.GetURL(), nil)
	if err != nil {
		client.debugLog(err)

		return nil, err
	}

	transport := http.DefaultTransport.(*http.Transport)

	resp, err := transport.RoundTrip(req)
	if err != nil {
		client.debugLog(err)

		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		client.debugLog("download status is not 200")

		return nil, err
	}

	return resp.Body, nil
}
