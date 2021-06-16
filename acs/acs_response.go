package acs

type Response interface{}

/*
	IsSuccess() bool
	GetHttpStatus() int
	GetHttpHeaders() map[string][]string
	GetHttpContentString() string
	GetHttpContentBytes() []byte
	GetOriginHttpResponse() *http.Response
	parseFromHttpResponse(httpResponse *http.Response) error
}
*/

type BaseResponse struct {
	httpStatus int
	// httpHeaders        map[string][]string
	// httpContentString  string
	// httpContentBytes   []byte
	// originHTTPResponse *http.Response
}

type CommonResponse struct {
	*BaseResponse
}

func NewCommonResponse() (response *CommonResponse) {
	return &CommonResponse{
		BaseResponse: &BaseResponse{},
	}
}

func (baseResponse *BaseResponse) GetHTTPStatus() int {
	return baseResponse.httpStatus
}

func (baseResponse *BaseResponse) IsSuccess() bool {
	if baseResponse.GetHTTPStatus() >= 200 && baseResponse.GetHTTPStatus() < 300 {
		return true
	}

	return false
}
