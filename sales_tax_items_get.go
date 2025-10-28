package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewSalesTaxItemsGetRequest() SalesTaxItemsGetRequest {
	r := SalesTaxItemsGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SalesTaxItemsGetRequest struct {
	client      *Client
	queryParams *SalesTaxItemsGetRequestQueryParams
	pathParams  *SalesTaxItemsGetRequestPathParams
	method      string
	headers     http.Header
	requestBody SalesTaxItemsGetRequestBody
}

func (r SalesTaxItemsGetRequest) NewQueryParams() *SalesTaxItemsGetRequestQueryParams {
	return &SalesTaxItemsGetRequestQueryParams{}
}

type SalesTaxItemsGetRequestQueryParams struct {
	SalesTaxItem string `schema:"SalesTaxItem,omitempty"`
	// PageSize       int    `schema:"PageSize"`
}

func (p SalesTaxItemsGetRequestQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *SalesTaxItemsGetRequest) QueryParams() *SalesTaxItemsGetRequestQueryParams {
	return r.queryParams
}

func (r SalesTaxItemsGetRequest) NewPathParams() *SalesTaxItemsGetRequestPathParams {
	return &SalesTaxItemsGetRequestPathParams{}
}

type SalesTaxItemsGetRequestPathParams struct {
}

func (p *SalesTaxItemsGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *SalesTaxItemsGetRequest) PathParams() *SalesTaxItemsGetRequestPathParams {
	return r.pathParams
}

func (r *SalesTaxItemsGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SalesTaxItemsGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SalesTaxItemsGetRequest) Method() string {
	return r.method
}

func (r SalesTaxItemsGetRequest) NewRequestBody() SalesTaxItemsGetRequestBody {
	return SalesTaxItemsGetRequestBody{}
}

type SalesTaxItemsGetRequestBody struct {
}

func (r *SalesTaxItemsGetRequest) RequestBody() *SalesTaxItemsGetRequestBody {
	return nil
}

func (r *SalesTaxItemsGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SalesTaxItemsGetRequest) SetRequestBody(body SalesTaxItemsGetRequestBody) {
	r.requestBody = body
}

func (r *SalesTaxItemsGetRequest) NewResponseBody() *SalesTaxItemsGetResponseBody {
	return &SalesTaxItemsGetResponseBody{}
}

type SalesTaxItemsGetResponseBody struct {
	Links Links `json:"links"`
	Count int   `json:"count"`
	HasMore bool  `json:"hasMore"`
	Items   []struct {
		Links []struct {
			Rel  string `json:"rel"`
			Href string `json:"href"`
		} `json:"links"`
		ID string `json:"id"`
	} `json:"items"`
	Offset       int `json:"offset"`
	TotalResults int `json:"totalResults"`
}

func (r *SalesTaxItemsGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/salesTaxItem", r.PathParams())
	return &u, err
}

func (r *SalesTaxItemsGetRequest) Do() (SalesTaxItemsGetResponseBody, error) {
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

