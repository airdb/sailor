package acs

import (
	"log"
	"net/url"
)

type Request interface {
	SetDebug()
	GetDebug() bool

	GetTimeout() int64

	SetDomain(domain string)
	GetDomain() string

	SetEndpoint(string)
	GetEndpoint() string

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

	SetUserAgent(string)
	GetUserAgent() string
}

type BaseRequest struct {
	debug      bool
	isInsecure bool
	method     string
	Scheme     string
	url        string
	domain     string
	// port       string
	endpoint string
	timeout  int64

	// Implement `curl --data-urlencode`
	values  url.Values
	headers map[string]string

	body interface{}
	/*
		RegionId       string
		ReadTimeout    time.Duration
		ConnectTimeout time.Duration

		userAgent map[string]string
		product   string
		version   string

		actionName string

		AcceptFormat string

		QueryParams map[string]string
		Headers     map[string]string
		FormParams  map[string]string
		Content     []byte

		locationServiceCode  string
		locationEndpointType string

		queries string

		stringToSign string
	*/
}

const (
	SchemeHTTP  = "http"
	SchemeHTTPS = "https"

	DefaultHTTPTimout = 60
)

const (
	UserAgentKey   = "User-Agent"
	UserAgentValue = "sailor-agent/1.0.0"

	ContentTypeKey  = "Content-Type"
	ContentTypeJSON = "application/json"
	ContentTypeXML  = "application/xml"
	ContentTypeRaw  = "application/octet-stream"
	ContentTypeForm = "application/x-www-form-urlencoded"
)

func (req *BaseRequest) SetDebug() {
	req.debug = true
}

func (req *BaseRequest) GetDebug() bool {
	return req.debug
}

func (req *BaseRequest) log(msg interface{}) {
	if req.debug {
		log.Println(msg)
	}
}

func (req *BaseRequest) SetMethod(method string) {
	req.method = method
}

func (req *BaseRequest) GetMethod() string {
	return req.method
}

func (req *BaseRequest) GetScheme() string {
	req.Scheme = SchemeHTTPS
	if req.isInsecure {
		req.Scheme = SchemeHTTP
	}

	return req.Scheme
}

func (req *BaseRequest) SetDomain(host string) {
	req.domain = host
}

func (req *BaseRequest) GetDomain() string {
	return req.domain
}

func (req *BaseRequest) GetHost() string {
	return req.domain
}

func (req *BaseRequest) SetEndpoint(endpoint string) {
	req.endpoint = endpoint
}

func (req *BaseRequest) GetEndpoint() string {
	return req.endpoint
}

func (req *BaseRequest) SetURL(url string) {
	req.url = url
}

func (req *BaseRequest) GetURL() string {
	if req.url != "" {
		return req.url
	}

	_url := url.URL{
		Scheme: req.GetScheme(),
		Host:   req.GetHost(),
		Path:   req.GetEndpoint(),
	}

	req.url = _url.String()
	req.log(req.url)

	return req.url
}

func (req *BaseRequest) GetTimeout() int64 {
	if req.timeout != 0 {
		return req.timeout
	}

	return DefaultHTTPTimout
}

func (req *BaseRequest) SetHeaders(headers map[string]string) {
	req.headers = headers
}

func (req *BaseRequest) GetHeaders() map[string]string {
	return req.headers
}

func (req *BaseRequest) SetValues(values url.Values) {
	req.values = values
}

func (req *BaseRequest) GetValues() url.Values {
	return req.values
}

func (req *BaseRequest) SetBody(body interface{}) {
	req.body = body
}

func (req *BaseRequest) GetBody() interface{} {
	return req.body
}

func (req *BaseRequest) SetUserAgent(ua string) {
	if req.headers == nil {
		header := make(map[string]string)
		req.headers = header
	}

	req.headers[UserAgentKey] = ua
}

func (req *BaseRequest) GetUserAgent() string {
	return req.headers[UserAgentKey]
}
