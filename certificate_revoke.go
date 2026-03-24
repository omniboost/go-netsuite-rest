package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewCertificateRevokeRequest() CertificateRevokeRequest {
	r := CertificateRevokeRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CertificateRevokeRequest struct {
	client      *Client
	queryParams *CertificateRevokeRequestQueryParams
	pathParams  *CertificateRevokeRequestPathParams
	method      string
	headers     http.Header
	requestBody CertificateRevokeRequestBody
}

func (r CertificateRevokeRequest) NewQueryParams() *CertificateRevokeRequestQueryParams {
	return &CertificateRevokeRequestQueryParams{}
}

type CertificateRevokeRequestQueryParams struct {
}

func (p CertificateRevokeRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CertificateRevokeRequest) QueryParams() *CertificateRevokeRequestQueryParams {
	return r.queryParams
}

func (r CertificateRevokeRequest) NewPathParams() *CertificateRevokeRequestPathParams {
	return &CertificateRevokeRequestPathParams{}
}

type CertificateRevokeRequestPathParams struct {
	ClientID      string `schema:"client_id"`
	CertificateID string `schema:"certificate_id"`
}

func (p *CertificateRevokeRequestPathParams) Params() map[string]string {
	return map[string]string{
		"client_id":      p.ClientID,
		"certificate_id": p.CertificateID,
	}
}

func (r *CertificateRevokeRequest) PathParams() *CertificateRevokeRequestPathParams {
	return r.pathParams
}

func (r *CertificateRevokeRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CertificateRevokeRequest) SetMethod(method string) {
	r.method = method
}

func (r *CertificateRevokeRequest) Method() string {
	return r.method
}

func (r CertificateRevokeRequest) NewRequestBody() CertificateRevokeRequestBody {
	return CertificateRevokeRequestBody{}
}

type CertificateRevokeRequestBody struct {
}

func (r *CertificateRevokeRequest) RequestBody() *CertificateRevokeRequestBody {
	return &r.requestBody
}

func (r *CertificateRevokeRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CertificateRevokeRequest) SetRequestBody(body CertificateRevokeRequestBody) {
	r.requestBody = body
}

func (r *CertificateRevokeRequest) NewResponseBody() *CertificateRevokeResponseBody {
	return &CertificateRevokeResponseBody{}
}

type CertificateRevokeResponseBody struct {
}

func (r *CertificateRevokeRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/auth/oauth2/v1/clients/{{.client_id}}/certificates/{{.certificate_id}}/revoke", r.PathParams())
	return &u, err
}

func (r *CertificateRevokeRequest) Do() (CertificateRevokeResponseBody, error) {
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
