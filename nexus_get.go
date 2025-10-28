package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewNexusGetRequest() NexusGetRequest {
	r := NexusGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type NexusGetRequest struct {
	client      *Client
	queryParams *NexusGetRequestQueryParams
	pathParams  *NexusGetRequestPathParams
	method      string
	headers     http.Header
	requestBody NexusGetRequestBody
}

func (r NexusGetRequest) NewQueryParams() *NexusGetRequestQueryParams {
	return &NexusGetRequestQueryParams{}
}

type NexusGetRequestQueryParams struct {
	Nexus string `schema:"Nexus,omitempty"`
	// PageSize       int    `schema:"PageSize"`
}

func (p NexusGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *NexusGetRequest) QueryParams() *NexusGetRequestQueryParams {
	return r.queryParams
}

func (r NexusGetRequest) NewPathParams() *NexusGetRequestPathParams {
	return &NexusGetRequestPathParams{}
}

type NexusGetRequestPathParams struct {
	ID string `schema:"id"`
}

func (p *NexusGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.ID,
	}
}

func (r *NexusGetRequest) PathParams() *NexusGetRequestPathParams {
	return r.pathParams
}

func (r *NexusGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *NexusGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *NexusGetRequest) Method() string {
	return r.method
}

func (r NexusGetRequest) NewRequestBody() NexusGetRequestBody {
	return NexusGetRequestBody{}
}

type NexusGetRequestBody struct {
}

func (r *NexusGetRequest) RequestBody() *NexusGetRequestBody {
	return nil
}

func (r *NexusGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *NexusGetRequest) SetRequestBody(body NexusGetRequestBody) {
	r.requestBody = body
}

func (r *NexusGetRequest) NewResponseBody() *NexusGetResponseBody {
	return &NexusGetResponseBody{}
}

type NexusGetResponseBody Nexus

func (r *NexusGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/nexus/{{.id}}", r.PathParams())
	return &u, err
}

func (r *NexusGetRequest) Do() (NexusGetResponseBody, error) {
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
