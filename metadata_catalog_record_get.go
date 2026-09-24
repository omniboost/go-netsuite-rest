package netsuite

import (
	"context"
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

// NewMetadataCatalogRecordGetRequest reads the JSON schema of one record type
// from the metadata catalog. The same URL also serves a swagger document; the
// schema is asked for by its media type, application/schema+json, which is
// why this request sets its own Accept header rather than the client's.
func (c *Client) NewMetadataCatalogRecordGetRequest() MetadataCatalogRecordGetRequest {
	r := MetadataCatalogRecordGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type MetadataCatalogRecordGetRequest struct {
	client      *Client
	queryParams *MetadataCatalogRecordGetRequestQueryParams
	pathParams  *MetadataCatalogRecordGetRequestPathParams
	method      string
	headers     http.Header
	requestBody MetadataCatalogRecordGetRequestBody
}

func (r MetadataCatalogRecordGetRequest) NewQueryParams() *MetadataCatalogRecordGetRequestQueryParams {
	return &MetadataCatalogRecordGetRequestQueryParams{}
}

type MetadataCatalogRecordGetRequestQueryParams struct {
}

func (p MetadataCatalogRecordGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *MetadataCatalogRecordGetRequest) QueryParams() *MetadataCatalogRecordGetRequestQueryParams {
	return r.queryParams
}

func (r MetadataCatalogRecordGetRequest) NewPathParams() *MetadataCatalogRecordGetRequestPathParams {
	return &MetadataCatalogRecordGetRequestPathParams{}
}

type MetadataCatalogRecordGetRequestPathParams struct {
	// Record names the record type, as the catalog listing spells it
	// (journalentry) or in camel case (journalEntry); NetSuite accepts both.
	Record string `schema:"record"`
}

func (p *MetadataCatalogRecordGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"record": p.Record,
	}
}

func (r *MetadataCatalogRecordGetRequest) PathParams() *MetadataCatalogRecordGetRequestPathParams {
	return r.pathParams
}

func (r *MetadataCatalogRecordGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *MetadataCatalogRecordGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *MetadataCatalogRecordGetRequest) Method() string {
	return r.method
}

func (r MetadataCatalogRecordGetRequest) NewRequestBody() MetadataCatalogRecordGetRequestBody {
	return MetadataCatalogRecordGetRequestBody{}
}

type MetadataCatalogRecordGetRequestBody struct {
}

func (r *MetadataCatalogRecordGetRequest) RequestBody() *MetadataCatalogRecordGetRequestBody {
	return nil
}

func (r *MetadataCatalogRecordGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *MetadataCatalogRecordGetRequest) SetRequestBody(body MetadataCatalogRecordGetRequestBody) {
	r.requestBody = body
}

func (r *MetadataCatalogRecordGetRequest) NewResponseBody() *MetadataCatalogRecordGetResponseBody {
	return &MetadataCatalogRecordGetResponseBody{}
}

// MetadataCatalogRecordGetResponseBody is the record type's schema itself:
// the endpoint returns the schema document with no envelope around it.
type MetadataCatalogRecordGetResponseBody struct {
	RecordSchema
}

func (r *MetadataCatalogRecordGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/metadata-catalog/{{.record}}", r.PathParams())
	return &u, err
}

func (r *MetadataCatalogRecordGetRequest) Do(ctx context.Context) (MetadataCatalogRecordGetResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(ctx, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// the same URL serves a swagger document under application/swagger+json;
	// the Accept header is what asks for the schema instead
	req.Header.Set("Accept", "application/schema+json")

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
