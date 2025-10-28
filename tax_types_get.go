package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewTaxTypesGetRequest() TaxTypesGetRequest {
	r := TaxTypesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type TaxTypesGetRequest struct {
	client      *Client
	queryParams *TaxTypesGetRequestQueryParams
	pathParams  *TaxTypesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody TaxTypesGetRequestBody
}

func (r TaxTypesGetRequest) NewQueryParams() *TaxTypesGetRequestQueryParams {
	return &TaxTypesGetRequestQueryParams{}
}

type TaxTypesGetRequestQueryParams struct {
	TaxType string `schema:"TaxType,omitempty"`
	// PageSize       int    `schema:"PageSize"`
}

func (p TaxTypesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *TaxTypesGetRequest) QueryParams() *TaxTypesGetRequestQueryParams {
	return r.queryParams
}

func (r TaxTypesGetRequest) NewPathParams() *TaxTypesGetRequestPathParams {
	return &TaxTypesGetRequestPathParams{}
}

type TaxTypesGetRequestPathParams struct {
}

func (p *TaxTypesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *TaxTypesGetRequest) PathParams() *TaxTypesGetRequestPathParams {
	return r.pathParams
}

func (r *TaxTypesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *TaxTypesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *TaxTypesGetRequest) Method() string {
	return r.method
}

func (r TaxTypesGetRequest) NewRequestBody() TaxTypesGetRequestBody {
	return TaxTypesGetRequestBody{}
}

type TaxTypesGetRequestBody struct {
}

func (r *TaxTypesGetRequest) RequestBody() *TaxTypesGetRequestBody {
	return nil
}

func (r *TaxTypesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *TaxTypesGetRequest) SetRequestBody(body TaxTypesGetRequestBody) {
	r.requestBody = body
}

func (r *TaxTypesGetRequest) NewResponseBody() *TaxTypeGetResponseBody {
	return &TaxTypeGetResponseBody{}
}

type TaxTypeGetResponseBody struct {
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

func (r *TaxTypesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/taxType", r.PathParams())
	return &u, err
}

func (r *TaxTypesGetRequest) Do() (TaxTypeGetResponseBody, error) {
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
