package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewSubsidiaryByIDGetRequest() SubsidiaryByIDGetRequest {
	r := SubsidiaryByIDGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SubsidiaryByIDGetRequest struct {
	client      *Client
	queryParams *SubsidiaryByIDGetRequestQueryParams
	pathParams  *SubsidiaryByIDGetRequestPathParams
	method      string
	headers     http.Header
	requestBody SubsidiaryByIDGetRequestBody
}

func (r SubsidiaryByIDGetRequest) NewQueryParams() *SubsidiaryByIDGetRequestQueryParams {
	return &SubsidiaryByIDGetRequestQueryParams{}
}

type SubsidiaryByIDGetRequestQueryParams struct {
	ExpandSubResources bool `schema:"expandSubResources,omitempty"`
}

func (p SubsidiaryByIDGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SubsidiaryByIDGetRequest) QueryParams() *SubsidiaryByIDGetRequestQueryParams {
	return r.queryParams
}

func (r SubsidiaryByIDGetRequest) NewPathParams() *SubsidiaryByIDGetRequestPathParams {
	return &SubsidiaryByIDGetRequestPathParams{}
}

type SubsidiaryByIDGetRequestPathParams struct {
	SubsidiaryID string `schema:"id"`
}

func (p *SubsidiaryByIDGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": p.SubsidiaryID,
	}
}

func (r *SubsidiaryByIDGetRequest) PathParams() *SubsidiaryByIDGetRequestPathParams {
	return r.pathParams
}

func (r *SubsidiaryByIDGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SubsidiaryByIDGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *SubsidiaryByIDGetRequest) Method() string {
	return r.method
}

func (r SubsidiaryByIDGetRequest) NewRequestBody() SubsidiaryByIDGetRequestBody {
	return SubsidiaryByIDGetRequestBody{}
}

type SubsidiaryByIDGetRequestBody struct {
}

func (r *SubsidiaryByIDGetRequest) RequestBody() *SubsidiaryByIDGetRequestBody {
	return nil
}

func (r *SubsidiaryByIDGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *SubsidiaryByIDGetRequest) SetRequestBody(body SubsidiaryByIDGetRequestBody) {
	r.requestBody = body
}

func (r *SubsidiaryByIDGetRequest) NewResponseBody() *SubsidiaryByIDGetResponseBody {
	return &SubsidiaryByIDGetResponseBody{}
}

type SubsidiaryByIDGetResponseBody struct {
	Subsidiary
}

func (r *SubsidiaryByIDGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/subsidiary/{{.id}}", r.PathParams())
	return &u, err
}

func (r *SubsidiaryByIDGetRequest) Do() (SubsidiaryByIDGetResponseBody, error) {
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
