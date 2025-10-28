package netsuite

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"

	"github.com/omniboost/go-netsuite-rest/utils"
)

func (c *Client) NewSuiteqlPostRequest() SuiteqlPostRequest {
	r := SuiteqlPostRequest{
		client:  c,
		method:  http.MethodPost,
		headers: http.Header{},
	}

	r.queryParams = r.NewQueryParams()
	r.pathParams = r.NewPathParams()
	r.requestBody = r.NewRequestBody()
	return r
}

type SuiteqlPostRequest struct {
	client      *Client
	queryParams *SuiteqlPostRequestQueryParams
	pathParams  *SuiteqlPostRequestPathParams
	method      string
	headers     http.Header
	requestBody SuiteqlPostRequestBody
}

func (r SuiteqlPostRequest) NewQueryParams() *SuiteqlPostRequestQueryParams {
	return &SuiteqlPostRequestQueryParams{}
}

type SuiteqlPostRequestQueryParams struct {
	Limit  int `schema:"limit,omitempty"`
	Offset int `schema:"offset,omitempty"`
}

func (p SuiteqlPostRequestQueryParams) ToURLValues() (url.Values, error) {
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

func (r *SuiteqlPostRequest) QueryParams() *SuiteqlPostRequestQueryParams {
	return r.queryParams
}

func (r SuiteqlPostRequest) NewPathParams() *SuiteqlPostRequestPathParams {
	return &SuiteqlPostRequestPathParams{}
}

type SuiteqlPostRequestPathParams struct {
	ID int `schema:"id"`
}

func (p *SuiteqlPostRequestPathParams) Params() map[string]string {
	return map[string]string{
		"id": strconv.Itoa(p.ID),
	}
}

func (r *SuiteqlPostRequest) PathParams() *SuiteqlPostRequestPathParams {
	return r.pathParams
}

func (r *SuiteqlPostRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *SuiteqlPostRequest) SetMethod(method string) {
	r.method = method
}

func (r *SuiteqlPostRequest) Method() string {
	return r.method
}

func (r SuiteqlPostRequest) NewRequestBody() SuiteqlPostRequestBody {
	return SuiteqlPostRequestBody{}
}

type SuiteqlPostRequestBody struct {
	Q string `json:"q"`
}

func (r *SuiteqlPostRequest) RequestBody() *SuiteqlPostRequestBody {
	return &r.requestBody
}

func (r *SuiteqlPostRequest) RequestBodyInterface() interface{} {
	return &r.requestBody
}

func (r *SuiteqlPostRequest) SetRequestBody(body SuiteqlPostRequestBody) {
	r.requestBody = body
}

func (r *SuiteqlPostRequest) NewResponseBody() *SuiteqlPostResponseBody {
	return &SuiteqlPostResponseBody{}
}

type SuiteqlPostResponseBody struct {
	Links        Links           `json:"links"`
	Count        int             `json:"count"`
	HasMore      bool            `json:"hasMore"`
	Items        json.RawMessage `json::"items"`
	Offset       int             `json:"offset"`
	TotalResults int             `json:"totalResults"`
}

func (r SuiteqlPostResponseBody) ToCustomers(client *Client) (Customers, error) {
	items := Customers{}

	reader := bytes.NewReader(r.Items)
	dec := json.NewDecoder(reader)
	if client.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err := dec.Decode(&items)
	return items, err
}

func (r *SuiteqlPostResponseBody) ToClassifications(client *Client) (Classifications, error) {
	items := Classifications{}

	reader := bytes.NewReader(r.Items)
	dec := json.NewDecoder(reader)
	if client.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err := dec.Decode(&items)
	return items, err
}

func (r *SuiteqlPostResponseBody) ToDepartments(client *Client) (Departments, error) {
	items := Departments{}

	reader := bytes.NewReader(r.Items)
	dec := json.NewDecoder(reader)
	if client.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err := dec.Decode(&items)
	return items, err
}

func (r *SuiteqlPostResponseBody) ToAccounts(client *Client) (Accounts, error) {
	items := Accounts{}
	err := json.Unmarshal(r.Items, &items)
	return items, err
}

func (r *SuiteqlPostResponseBody) ToAddresses(client *Client) (Addresses, error) {
	items := Addresses{}
	err := json.Unmarshal(r.Items, &items)
	return items, err
}

func (r *SuiteqlPostResponseBody) ToLocations(client *Client) (Locations, error) {
	items := Locations{}

	reader := bytes.NewReader(r.Items)
	dec := json.NewDecoder(reader)
	if client.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err := dec.Decode(&items)
	return items, err
}

func (r *SuiteqlPostResponseBody) ToCreditMemoTransactions(client *Client) (CreditMemoTransactions, error) {
	items := CreditMemoTransactions{}

	reader := bytes.NewReader(r.Items)
	dec := json.NewDecoder(reader)
	if client.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err := dec.Decode(&items)
	return items, err
}

func (r *SuiteqlPostResponseBody) ToInvoiceTransactions(client *Client) (InvoiceTransactions, error) {
	items := InvoiceTransactions{}

	reader := bytes.NewReader(r.Items)
	dec := json.NewDecoder(reader)
	if client.disallowUnknownFields {
		dec.DisallowUnknownFields()
	}
	err := dec.Decode(&items)
	return items, err
}

func (r *SuiteqlPostRequest) URL() (*url.URL, error) {
	u, err := r.client.GetEndpointURL("/query/v1/suiteql", r.PathParams())
	return &u, err
}

func (r *SuiteqlPostRequest) Do() (SuiteqlPostResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	req.Header.Add("Prefer", "transient")
	req.Header.Del("Content-Language")

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, false)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
