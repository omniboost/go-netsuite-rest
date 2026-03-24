package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewCertificatesGetRequest() CertificatesGetRequest {
	r := CertificatesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CertificatesGetRequest struct {
	client      *Client
	queryParams *CertificatesGetRequestQueryParams
	pathParams  *CertificatesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody CertificatesGetRequestBody
}

func (r CertificatesGetRequest) NewQueryParams() *CertificatesGetRequestQueryParams {
	return &CertificatesGetRequestQueryParams{}
}

type CertificatesGetRequestQueryParams struct {
	Q      string `schema:"q,omitempty"`
	Limit  int    `schema:"limit,omitempty"`
	Offset int    `schema:"offset,omitempty"`
}

func (p CertificatesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CertificatesGetRequest) QueryParams() *CertificatesGetRequestQueryParams {
	return r.queryParams
}

func (r CertificatesGetRequest) NewPathParams() *CertificatesGetRequestPathParams {
	return &CertificatesGetRequestPathParams{}
}

type CertificatesGetRequestPathParams struct {
	ClientID string `schema:"client_id,omitempty"`
}

func (p *CertificatesGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"client_id": p.ClientID,
	}
}

func (r *CertificatesGetRequest) PathParams() *CertificatesGetRequestPathParams {
	return r.pathParams
}

func (r *CertificatesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CertificatesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *CertificatesGetRequest) Method() string {
	return r.method
}

func (r CertificatesGetRequest) NewRequestBody() CertificatesGetRequestBody {
	return CertificatesGetRequestBody{}
}

type CertificatesGetRequestBody struct {
}

func (r *CertificatesGetRequest) RequestBody() *CertificatesGetRequestBody {
	return nil
}

func (r *CertificatesGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *CertificatesGetRequest) SetRequestBody(body CertificatesGetRequestBody) {
	r.requestBody = body
}

func (r *CertificatesGetRequest) NewResponseBody() *CertificatesGetResponseBody {
	return &CertificatesGetResponseBody{}
}

type CertificatesGetResponseBody Certificates

func (r *CertificatesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/auth/oauth2/v1/clients/{{.client_id}}/certificates", r.PathParams())
	return &u, err
}

func (r *CertificatesGetRequest) Do() (CertificatesGetResponseBody, error) {
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
