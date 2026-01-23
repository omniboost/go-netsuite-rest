package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewInvoicePatchRequest() InvoicePatchRequest {
	r := InvoicePatchRequest{
		client:  c,
		method:  http.MethodPatch,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type InvoicePatchRequest struct {
	client      *Client
	queryParams *InvoicePatchRequestQueryParams
	pathParams  *InvoicePatchRequestPathParams
	method      string
	headers     http.Header
	requestBody InvoicePatchRequestBody
}

func (r InvoicePatchRequest) NewQueryParams() *InvoicePatchRequestQueryParams {
	return &InvoicePatchRequestQueryParams{}
}

type InvoicePatchRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p InvoicePatchRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *InvoicePatchRequest) QueryParams() *InvoicePatchRequestQueryParams {
	return r.queryParams
}

func (r InvoicePatchRequest) NewPathParams() *InvoicePatchRequestPathParams {
	return &InvoicePatchRequestPathParams{}
}

type InvoicePatchRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *InvoicePatchRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *InvoicePatchRequest) PathParams() *InvoicePatchRequestPathParams {
	return r.pathParams
}

func (r *InvoicePatchRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *InvoicePatchRequest) SetMethod(method string) {
	r.method = method
}

func (r *InvoicePatchRequest) Method() string {
	return r.method
}

func (r InvoicePatchRequest) NewRequestBody() InvoicePatchRequestBody {
	return InvoicePatchRequestBody{}
}

type InvoicePatchRequestBody struct {
	Invoice
}

func (r *InvoicePatchRequest) RequestBody() *InvoicePatchRequestBody {
	return &r.requestBody
}

func (r *InvoicePatchRequest) RequestBodyInterface() any {
	return &r.requestBody
}

func (r *InvoicePatchRequest) SetRequestBody(body InvoicePatchRequestBody) {
	r.requestBody = body
}

func (r *InvoicePatchRequest) NewResponseBody() *InvoicePatchResponseBody {
	return &InvoicePatchResponseBody{}
}

type InvoicePatchResponseBody struct {
}

func (r *InvoicePatchRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/invoice/{{.id}}", r.PathParams())
	return &u, err
}

func (r *InvoicePatchRequest) Do() (InvoicePatchResponseBody, error) {
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
