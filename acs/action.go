package acs

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

func (client *Client) DoAction(req Request, resp Response) error {
	if req.GetURL() == "" {
		return errors.New("API url is null, please check")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(req.GetTimeout()))
	defer cancel()

	switch req.GetMethod() {
	case http.MethodGet:
		return client.HTTPRequestGet(ctx, req, resp)
	case http.MethodPost, http.MethodPatch:
		return client.HTTPRequestPost(ctx, req, resp)
	default:
		req.SetMethod(http.MethodGet)

		return client.HTTPRequestGet(ctx, req, resp)
	}
}

func (client *Client) HTTPRequestGet(ctx context.Context, req Request, resp Response) error {
	/*
		var r *strings.Reader

		if requestInterface.GetValues() != nil {
			r = strings.NewReader(requestInterface.GetValues().Encode())
		}
		log.Println(r)
	*/
	request, err := http.NewRequestWithContext(ctx, req.GetMethod(), req.GetURL(), nil)
	if err != nil {
		return err
	}

	if req.GetValues() != nil {
		request.Header.Add(ContentTypeKey, ContentTypeForm)
	}

	// The RawQuery should binding on `url`.  Like follow:
	// Query string `url:"q"`
	body, err := query.Values(req.GetBody())
	if err != nil {
		return err
	}

	if len(body) > 0 {
		request.URL.RawQuery = body.Encode()
	}

	out, err := DoRequest(request, req)
	if err != nil {
		return err
	}

	defer out.Body.Close()

	switch out.StatusCode {
	case http.StatusOK:
		return json.NewDecoder(out.Body).Decode(&resp)
	default:
		return errors.Errorf("http status code is not 200.")
	}
}

func (client *Client) HTTPRequestPost(ctx context.Context, req Request, resp Response) error {
	payload, err := json.Marshal(req.GetBody())
	if err != nil {
		return err
	}

	request, err := http.NewRequestWithContext(ctx, req.GetMethod(), req.GetURL(), bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	request.Header.Add(ContentTypeKey, ContentTypeJSON)

	out, err := DoRequest(request, req)
	if err != nil {
		return err
	}

	defer out.Body.Close()

	switch out.StatusCode {
	case http.StatusOK:
		return json.NewDecoder(out.Body).Decode(&resp)
	default:
		return errors.Errorf("http status code is not 200.")
	}
}

func DoRequest(request *http.Request, req Request) (*http.Response, error) {
	for headerKey, headerValue := range req.GetHeaders() {
		request.Header.Set(headerKey, headerValue)
	}

	headers := req.GetHeaders()
	if headers[UserAgentKey] == "" {
		request.Header.Set(UserAgentKey, UserAgentValue)
	}

	_client := http.DefaultClient

	return _client.Do(request)
}

func (client *Client) Downloader(req Request) (io.ReadCloser, error) {
	_httpRequest, err := http.NewRequestWithContext(context.Background(), http.MethodGet, req.GetURL(), nil)
	if err != nil {
		return nil, err
	}

	transport := http.DefaultTransport.(*http.Transport)

	resp, err := transport.RoundTrip(_httpRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	return resp.Body, nil
}
