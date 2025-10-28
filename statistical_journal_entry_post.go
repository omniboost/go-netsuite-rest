package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewStatisticalJournalEntryPostRequest() StatisticalJournalEntryPostRequest {
	r := StatisticalJournalEntryPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type StatisticalJournalEntryPostRequest struct {
	client      *Client
	queryParams *StatisticalJournalEntryPostRequestQueryParams
	pathParams  *StatisticalJournalEntryPostRequestPathParams
	method      string
	headers     http.Header
	requestBody StatisticalJournalEntryPostRequestBody
}

func (r StatisticalJournalEntryPostRequest) NewQueryParams() *StatisticalJournalEntryPostRequestQueryParams {
	return &StatisticalJournalEntryPostRequestQueryParams{}
}

type StatisticalJournalEntryPostRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p StatisticalJournalEntryPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *StatisticalJournalEntryPostRequest) QueryParams() *StatisticalJournalEntryPostRequestQueryParams {
	return r.queryParams
}

func (r StatisticalJournalEntryPostRequest) NewPathParams() *StatisticalJournalEntryPostRequestPathParams {
	return &StatisticalJournalEntryPostRequestPathParams{}
}

type StatisticalJournalEntryPostRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *StatisticalJournalEntryPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *StatisticalJournalEntryPostRequest) PathParams() *StatisticalJournalEntryPostRequestPathParams {
	return r.pathParams
}

func (r *StatisticalJournalEntryPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *StatisticalJournalEntryPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *StatisticalJournalEntryPostRequest) Method() string {
	return r.method
}

func (r StatisticalJournalEntryPostRequest) NewRequestBody() StatisticalJournalEntryPostRequestBody {
	return StatisticalJournalEntryPostRequestBody{}
}

type StatisticalJournalEntryPostRequestBody struct {
	StatisticalJournalEntry
}

func (r *StatisticalJournalEntryPostRequest) RequestBody() *StatisticalJournalEntryPostRequestBody {
	return &r.requestBody
}

func (r *StatisticalJournalEntryPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *StatisticalJournalEntryPostRequest) SetRequestBody(body StatisticalJournalEntryPostRequestBody) {
	r.requestBody = body
}

func (r *StatisticalJournalEntryPostRequest) NewResponseBody() *StatisticalJournalEntryPostResponseBody {
	return &StatisticalJournalEntryPostResponseBody{}
}

type StatisticalJournalEntryPostResponseBody struct {
}

func (r *StatisticalJournalEntryPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/statisticalJournalEntry", r.PathParams())
	return &u, err
}

func (r *StatisticalJournalEntryPostRequest) Do() (StatisticalJournalEntryPostResponseBody, error) {
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
