package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewSalesTaxItemGetRequest() SalesTaxItemGetRequest {
	r := SalesTaxItemGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesTaxItemGetRequest struct {
	client      *Client
	queryParams *SalesTaxItemGetRequestQueryParams
	pathParams  *SalesTaxItemGetRequestPathParams
	method      string
	headers     http.Header
	requestBody SalesTaxItemGetRequestBody
}

func (r SalesTaxItemGetRequest) NewQueryParams() *SalesTaxItemGetRequestQueryParams {
	return &SalesTaxItemGetRequestQueryParams{}
}

type SalesTaxItemGetRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p SalesTaxItemGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SalesTaxItemGetRequest) QueryParams() *SalesTaxItemGetRequestQueryParams {
	return r.queryParams
}

func (r SalesTaxItemGetRequest) NewPathParams() *SalesTaxItemGetRequestPathParams {
	return &SalesTaxItemGetRequestPathParams{}
}

type SalesTaxItemGetRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *SalesTaxItemGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *SalesTaxItemGetRequest) PathParams() *SalesTaxItemGetRequestPathParams {
	return r.pathParams
}

func (r *SalesTaxItemGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesTaxItemGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesTaxItemGetRequest) Method() string {
	return r.method
}

func (r SalesTaxItemGetRequest) NewRequestBody() SalesTaxItemGetRequestBody {
	return SalesTaxItemGetRequestBody{}
}

type SalesTaxItemGetRequestBody struct {
}

func (r *SalesTaxItemGetRequest) RequestBody() *SalesTaxItemGetRequestBody {
	return nil
}

func (r *SalesTaxItemGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesTaxItemGetRequest) SetRequestBody(body SalesTaxItemGetRequestBody) {
	r.requestBody = body
}

func (r *SalesTaxItemGetRequest) NewResponseBody() *SalesTaxItemGetResponseBody {
	return &SalesTaxItemGetResponseBody{}
}

type SalesTaxItemGetResponseBody struct {
	Links Links `json:"links"`
	SalesTaxItem
}

func (r *SalesTaxItemGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/salesTaxItem/{{.id}}", r.PathParams())
	return &u, err
}

func (r *SalesTaxItemGetRequest) Do() (SalesTaxItemGetResponseBody, error) {
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

