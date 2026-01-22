package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewLocationGetRequest() LocationGetRequest {
	r := LocationGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type LocationGetRequest struct {
	client      *Client
	queryParams *LocationGetRequestQueryParams
	pathParams  *LocationGetRequestPathParams
	method      string
	headers     http.Header
	requestBody LocationGetRequestBody
}

func (r LocationGetRequest) NewQueryParams() *LocationGetRequestQueryParams {
	return &LocationGetRequestQueryParams{}
}

type LocationGetRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p LocationGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *LocationGetRequest) QueryParams() *LocationGetRequestQueryParams {
	return r.queryParams
}

func (r LocationGetRequest) NewPathParams() *LocationGetRequestPathParams {
	return &LocationGetRequestPathParams{}
}

type LocationGetRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *LocationGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *LocationGetRequest) PathParams() *LocationGetRequestPathParams {
	return r.pathParams
}

func (r *LocationGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *LocationGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *LocationGetRequest) Method() string {
	return r.method
}

func (r LocationGetRequest) NewRequestBody() LocationGetRequestBody {
	return LocationGetRequestBody{}
}

type LocationGetRequestBody struct {
}

func (r *LocationGetRequest) RequestBody() *LocationGetRequestBody {
	return nil
}

func (r *LocationGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *LocationGetRequest) SetRequestBody(body LocationGetRequestBody) {
	r.requestBody = body
}

func (r *LocationGetRequest) NewResponseBody() *LocationGetResponseBody {
	return &LocationGetResponseBody{}
}

type LocationGetResponseBody struct {
	Links Links `json:"links"`
	Location
}

func (r *LocationGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/location/{{.id}}", r.PathParams())
	return &u, err
}

func (r *LocationGetRequest) Do() (LocationGetResponseBody, error) {
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
