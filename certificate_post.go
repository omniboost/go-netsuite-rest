package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewCertificatePostRequest() CertificatePostRequest {
	r := CertificatePostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type CertificatePostRequest struct {
	client      *Client
	queryParams *CertificatePostRequestQueryParams
	pathParams  *CertificatePostRequestPathParams
	method      string
	headers     http.Header
	requestBody CertificatePostRequestBody
}

func (r CertificatePostRequest) NewQueryParams() *CertificatePostRequestQueryParams {
	return &CertificatePostRequestQueryParams{}
}

type CertificatePostRequestQueryParams struct {
}

func (p CertificatePostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *CertificatePostRequest) QueryParams() *CertificatePostRequestQueryParams {
	return r.queryParams
}

func (r CertificatePostRequest) NewPathParams() *CertificatePostRequestPathParams {
	return &CertificatePostRequestPathParams{}
}

type CertificatePostRequestPathParams struct {
	ClientID string `schema:"client_id"`
}

func (p *CertificatePostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"client_id": p.ClientID,
	}
}

func (r *CertificatePostRequest) PathParams() *CertificatePostRequestPathParams {
	return r.pathParams
}

func (r *CertificatePostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *CertificatePostRequest) SetMethod(method string) {
	r.method = method
}

func (r *CertificatePostRequest) Method() string {
	return r.method
}

func (r CertificatePostRequest) NewRequestBody() CertificatePostRequestBody {
	return CertificatePostRequestBody{}
}

type CertificatePostRequestBody struct {
	// The value of the parameter is the certificate you want to set up.
	FileContent string `json:"fileContent"`
	// The value of the parameter is the ID of the role for which you want to activate the certificate.
	Role        int    `json:"role"`
	// The value of the parameter is the ID of the entity for which you want to activate the certificate.
	Entity      int    `json:"entity"`
}

func (r *CertificatePostRequest) RequestBody() *CertificatePostRequestBody {
	return &r.requestBody
}

func (r *CertificatePostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *CertificatePostRequest) SetRequestBody(body CertificatePostRequestBody) {
	r.requestBody = body
}

func (r *CertificatePostRequest) NewResponseBody() *CertificatePostResponseBody {
	return &CertificatePostResponseBody{}
}

type CertificatePostResponseBody struct {
}

func (r *CertificatePostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/auth/oauth2/v1/clients/{{.client_id}}/certificates", r.PathParams())
	return &u, err
}

func (r *CertificatePostRequest) Do() (CertificatePostResponseBody, error) {
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
