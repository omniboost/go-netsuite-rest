package netsuite

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

// NewMetadataCatalogGetRequest lists every record type the metadata catalog
// describes for the account, standard and custom alike. The names it returns
// are in lower case (journalentry), which the per-record request accepts as
// well as camel case (journalEntry).
func (c *Client) NewMetadataCatalogGetRequest() MetadataCatalogGetRequest {
	r := MetadataCatalogGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MetadataCatalogGetRequest struct {
	client      *Client
	queryParams *MetadataCatalogGetRequestQueryParams
	pathParams  *MetadataCatalogGetRequestPathParams
	method      string
	headers     http.Header
	requestBody MetadataCatalogGetRequestBody
}

func (r MetadataCatalogGetRequest) NewQueryParams() *MetadataCatalogGetRequestQueryParams {
	return &MetadataCatalogGetRequestQueryParams{}
}

type MetadataCatalogGetRequestQueryParams struct {
	// Select narrows the listing to the named record types, comma separated.
	Select string `schema:"select,omitempty"`
}

func (p MetadataCatalogGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MetadataCatalogGetRequest) QueryParams() *MetadataCatalogGetRequestQueryParams {
	return r.queryParams
}

func (r MetadataCatalogGetRequest) NewPathParams() *MetadataCatalogGetRequestPathParams {
	return &MetadataCatalogGetRequestPathParams{}
}

type MetadataCatalogGetRequestPathParams struct {
}

func (p *MetadataCatalogGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *MetadataCatalogGetRequest) PathParams() *MetadataCatalogGetRequestPathParams {
	return r.pathParams
}

func (r *MetadataCatalogGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MetadataCatalogGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MetadataCatalogGetRequest) Method() string {
	return r.method
}

func (r MetadataCatalogGetRequest) NewRequestBody() MetadataCatalogGetRequestBody {
	return MetadataCatalogGetRequestBody{}
}

type MetadataCatalogGetRequestBody struct {
}

func (r *MetadataCatalogGetRequest) RequestBody() *MetadataCatalogGetRequestBody {
	return nil
}

func (r *MetadataCatalogGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MetadataCatalogGetRequest) SetRequestBody(body MetadataCatalogGetRequestBody) {
	r.requestBody = body
}

func (r *MetadataCatalogGetRequest) NewResponseBody() *MetadataCatalogGetResponseBody {
	return &MetadataCatalogGetResponseBody{}
}

type MetadataCatalogGetResponseBody struct {
	Items []MetadataCatalogItem `json:"items"`
	Links Links                 `json:"links"`
}

func (r *MetadataCatalogGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/metadata-catalog", r.PathParams())
	return &u, err
}

func (r *MetadataCatalogGetRequest) Do(ctx context.Context) (MetadataCatalogGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
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
