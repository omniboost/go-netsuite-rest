package netsuite

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewStatisticalJournalEntryGetRequest() StatisticalJournalEntryGetRequest {
	r := StatisticalJournalEntryGetRequest{
		client:  c,
		method:  http.MethodGet,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type StatisticalJournalEntryGetRequest struct {
	client      *Client
	queryParams *StatisticalJournalEntryGetRequestQueryParams
	pathParams  *StatisticalJournalEntryGetRequestPathParams
	method      string
	headers     http.Header
	requestBody StatisticalJournalEntryGetRequestBody
}

func (r StatisticalJournalEntryGetRequest) NewQueryParams() *StatisticalJournalEntryGetRequestQueryParams {
	return &StatisticalJournalEntryGetRequestQueryParams{}
}

type StatisticalJournalEntryGetRequestQueryParams struct {
	Fields             Fields `schema:"fields,omitempty"`
	ExpandSubResources bool   `schema:"expandSubResources,omitempty"`
}

func (p StatisticalJournalEntryGetRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *StatisticalJournalEntryGetRequest) QueryParams() *StatisticalJournalEntryGetRequestQueryParams {
	return r.queryParams
}

func (r StatisticalJournalEntryGetRequest) NewPathParams() *StatisticalJournalEntryGetRequestPathParams {
	return &StatisticalJournalEntryGetRequestPathParams{}
}

type StatisticalJournalEntryGetRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *StatisticalJournalEntryGetRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *StatisticalJournalEntryGetRequest) PathParams() *StatisticalJournalEntryGetRequestPathParams {
	return r.pathParams
}

func (r *StatisticalJournalEntryGetRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *StatisticalJournalEntryGetRequest) SetMethod(method string) {
	r.method = method
}

func (r *StatisticalJournalEntryGetRequest) Method() string {
	return r.method
}

func (r StatisticalJournalEntryGetRequest) NewRequestBody() StatisticalJournalEntryGetRequestBody {
	return StatisticalJournalEntryGetRequestBody{}
}

type StatisticalJournalEntryGetRequestBody struct {
}

func (r *StatisticalJournalEntryGetRequest) RequestBody() *StatisticalJournalEntryGetRequestBody {
	return nil
}

func (r *StatisticalJournalEntryGetRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *StatisticalJournalEntryGetRequest) SetRequestBody(body StatisticalJournalEntryGetRequestBody) {
	r.requestBody = body
}

func (r *StatisticalJournalEntryGetRequest) NewResponseBody() *StatisticalJournalEntryGetResponseBody {
	return &StatisticalJournalEntryGetResponseBody{}
}

type StatisticalJournalEntryGetResponseBody struct {
	Links Links `json:"links"`
	JournalEntry
}

func (r *StatisticalJournalEntryGetRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/record/v1/statisticalJournalEntry/{{.id}}", r.PathParams())
	return &u, err
}

func (r *StatisticalJournalEntryGetRequest) Do() (StatisticalJournalEntryGetResponseBody, error) {
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

