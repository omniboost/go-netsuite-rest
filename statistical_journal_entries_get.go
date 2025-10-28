package netsuite

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewStatisticalJournalEntriesGetRequest() StatisticalJournalEntriesGetRequest {
	r := StatisticalJournalEntriesGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type StatisticalJournalEntriesGetRequest struct {
	client      *Client
	queryParams *StatisticalJournalEntriesGetRequestQueryParams
	pathParams  *StatisticalJournalEntriesGetRequestPathParams
	method      string
	headers     http.Header
	requestBody StatisticalJournalEntriesGetRequestBody
}

func (r StatisticalJournalEntriesGetRequest) NewQueryParams() *StatisticalJournalEntriesGetRequestQueryParams {
	return &StatisticalJournalEntriesGetRequestQueryParams{}
}

type StatisticalJournalEntriesGetRequestQueryParams struct {
	Q      string `schema:"q,omitempty"`
	Limit  int    `schema:"limit,omitempty"`
	Offset int    `schema:"offset,omitempty"`
}

func (p StatisticalJournalEntriesGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *StatisticalJournalEntriesGetRequest) QueryParams() *StatisticalJournalEntriesGetRequestQueryParams {
	return r.queryParams
}

func (r StatisticalJournalEntriesGetRequest) NewPathParams() *StatisticalJournalEntriesGetRequestPathParams {
	return &StatisticalJournalEntriesGetRequestPathParams{}
}

type StatisticalJournalEntriesGetRequestPathParams struct {
}

func (p *StatisticalJournalEntriesGetRequestPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *StatisticalJournalEntriesGetRequest) PathParams() *StatisticalJournalEntriesGetRequestPathParams {
	return r.pathParams
}

func (r *StatisticalJournalEntriesGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *StatisticalJournalEntriesGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *StatisticalJournalEntriesGetRequest) Method() string {
	return r.method
}

func (r StatisticalJournalEntriesGetRequest) NewRequestBody() StatisticalJournalEntriesGetRequestBody {
	return StatisticalJournalEntriesGetRequestBody{}
}

type StatisticalJournalEntriesGetRequestBody struct {
}

func (r *StatisticalJournalEntriesGetRequest) RequestBody() *StatisticalJournalEntriesGetRequestBody {
	return nil
}

func (r *StatisticalJournalEntriesGetRequest) RequestBodyInterface() any {
	return nil
}

func (r *StatisticalJournalEntriesGetRequest) SetRequestBody(body StatisticalJournalEntriesGetRequestBody) {
	r.requestBody = body
}

func (r *StatisticalJournalEntriesGetRequest) NewResponseBody() *StatisticalJournalEntriesGetResponseBody {
	return &StatisticalJournalEntriesGetResponseBody{}
}

type StatisticalJournalEntriesGetResponseBody StatisticalJournalEntryCollection

func (r *StatisticalJournalEntriesGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/statisticalJournalEntry", r.PathParams())
	return &u, err
}

func (r *StatisticalJournalEntriesGetRequest) Do() (StatisticalJournalEntriesGetResponseBody, error) {
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
