package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewNexusesGetRequest() NexusesGetRequest {
	r := NexusesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type NexusesGetRequest struct {
	client      *Client
	queryParams *NexusesGetRequestQueryParams
	pathParams  *NexusesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody NexusesGetRequestBody
}

func (r NexusesGetRequest) NewQueryParams() *NexusesGetRequestQueryParams {
	return &NexusesGetRequestQueryParams{}
}

type NexusesGetRequestQueryParams struct {
	Q      string `schema:"q,omitempty"`
	Limit  int    `schema:"limit,omitempty"`
	Offset int    `schema:"offset,omitempty"`
}

func (p NexusesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *NexusesGetRequest) QueryParams() *NexusesGetRequestQueryParams {
	return r.queryParams
}

func (r NexusesGetRequest) NewPathParams() *NexusesGetRequestPathParams {
	return &NexusesGetRequestPathParams{}
}

type NexusesGetRequestPathParams struct {
}

func (p *NexusesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *NexusesGetRequest) PathParams() *NexusesGetRequestPathParams {
	return r.pathParams
}

func (r *NexusesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *NexusesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *NexusesGetRequest) Method() string {
	return r.method
}

func (r NexusesGetRequest) NewRequestBody() NexusesGetRequestBody {
	return NexusesGetRequestBody{}
}

type NexusesGetRequestBody struct {
}

func (r *NexusesGetRequest) RequestBody() *NexusesGetRequestBody {
	return nil
}

func (r *NexusesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *NexusesGetRequest) SetRequestBody(body NexusesGetRequestBody) {
	r.requestBody = body
}

func (r *NexusesGetRequest) NewResponseBody() *NexusesGetResponseBody {
	return &NexusesGetResponseBody{}
}

type NexusesGetResponseBody NexusCollection

func (r *NexusesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/nexus", r.PathParams())
	return &u, err
}

func (r *NexusesGetRequest) Do() (NexusesGetResponseBody, error) {
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

