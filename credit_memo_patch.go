package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewCreditMemoPatchRequest() CreditMemoPatchRequest {
	r := CreditMemoPatchRequest{
		client:  c,
		method:  http.MethodPatch,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CreditMemoPatchRequest struct {
	client      *Client
	queryParams *CreditMemoPatchRequestQueryParams
	pathParams  *CreditMemoPatchRequestPathParams
	method      string
	headers     http.Header
	requestBody CreditMemoPatchRequestBody
}

func (r CreditMemoPatchRequest) NewQueryParams() *CreditMemoPatchRequestQueryParams {
	return &CreditMemoPatchRequestQueryParams{}
}

type CreditMemoPatchRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p CreditMemoPatchRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CreditMemoPatchRequest) QueryParams() *CreditMemoPatchRequestQueryParams {
	return r.queryParams
}

func (r CreditMemoPatchRequest) NewPathParams() *CreditMemoPatchRequestPathParams {
	return &CreditMemoPatchRequestPathParams{}
}

type CreditMemoPatchRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *CreditMemoPatchRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *CreditMemoPatchRequest) PathParams() *CreditMemoPatchRequestPathParams {
	return r.pathParams
}

func (r *CreditMemoPatchRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CreditMemoPatchRequest) SetMethod(method string) {
	r.method = method
}

func (r *CreditMemoPatchRequest) Method() string {
	return r.method
}

func (r CreditMemoPatchRequest) NewRequestBody() CreditMemoPatchRequestBody {
	return CreditMemoPatchRequestBody{}
}

type CreditMemoPatchRequestBody struct {
	CreditMemo
}

func (r *CreditMemoPatchRequest) RequestBody() *CreditMemoPatchRequestBody {
	return &r.requestBody
}

func (r *CreditMemoPatchRequest) RequestBodyInterface() any {
	return &r.requestBody
}

func (r *CreditMemoPatchRequest) SetRequestBody(body CreditMemoPatchRequestBody) {
	r.requestBody = body
}

func (r *CreditMemoPatchRequest) NewResponseBody() *CreditMemoPatchResponseBody {
	return &CreditMemoPatchResponseBody{}
}

type CreditMemoPatchResponseBody struct {
}

func (r *CreditMemoPatchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/creditMemo/{{.id}}", r.PathParams())
	return &u, err
}

func (r *CreditMemoPatchRequest) Do() (CreditMemoPatchResponseBody, error) {
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

