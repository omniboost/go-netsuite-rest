package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewCreditMemoPostRequest() CreditMemoPostRequest {
	r := CreditMemoPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreditMemoPostRequest struct {
	client      *Client
	queryParams *CreditMemoPostRequestQueryParams
	pathParams  *CreditMemoPostRequestPathParams
	method      string
	headers     http.Header
	requestBody CreditMemoPostRequestBody
}

func (r CreditMemoPostRequest) NewQueryParams() *CreditMemoPostRequestQueryParams {
	return &CreditMemoPostRequestQueryParams{}
}

type CreditMemoPostRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p CreditMemoPostRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(Fields{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *CreditMemoPostRequest) QueryParams() *CreditMemoPostRequestQueryParams {
	return r.queryParams
}

func (r CreditMemoPostRequest) NewPathParams() *CreditMemoPostRequestPathParams {
	return &CreditMemoPostRequestPathParams{}
}

type CreditMemoPostRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *CreditMemoPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *CreditMemoPostRequest) PathParams() *CreditMemoPostRequestPathParams {
	return r.pathParams
}

func (r *CreditMemoPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CreditMemoPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreditMemoPostRequest) Method() string {
	return r.method
}

func (r CreditMemoPostRequest) NewRequestBody() CreditMemoPostRequestBody {
	return CreditMemoPostRequestBody{}
}

type CreditMemoPostRequestBody struct {
	CreditMemo
}

func (r *CreditMemoPostRequest) RequestBody() *CreditMemoPostRequestBody {
	return &r.requestBody
}

func (r *CreditMemoPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CreditMemoPostRequest) SetRequestBody(body CreditMemoPostRequestBody) {
	r.requestBody = body
}

func (r *CreditMemoPostRequest) NewResponseBody() *CreditMemoPostResponseBody {
	return &CreditMemoPostResponseBody{}
}

type CreditMemoPostResponseBody struct {
}

func (r *CreditMemoPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/creditMemo", r.PathParams())
	return &u, err
}

func (r *CreditMemoPostRequest) Do() (CreditMemoPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}

