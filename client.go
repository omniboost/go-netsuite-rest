package netsuite

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/pkg/errors"
)

const (
	libraryVersion = "0.0.1"
	userAgent      = "go-netsuite-rest/" + libraryVersion
	mediaType      = "application/json"
	charset        = "utf-8"
)

var (
	BaseURL string = "https://{{.account_id}}.suitetalk.api.netsuite.com/services/rest"
)

// NewClient returns a new Exact Globe Client client
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	client := &Client{}

	client.SetHTTPClient(httpClient)
	client.SetBaseURL(BaseURL)
	client.SetDebug(false)
	client.SetUserAgent(userAgent)
	client.SetMediaType(mediaType)
	client.SetCharset(charset)

	return client
}

// Client manages communication with Exact Globe Client
type Client struct {
	// HTTP client used to communicate with the Client.
	http *http.Client

	debug   bool
	baseURL string

	// credentials
	companyID       string
	contentLanguage string

	// token based auth credentials
	useTokenAuth bool
	clientID     string
	clientSecret string
	tokenID      string
	tokenSecret  string
	// accountID    string

	// User agent for client
	userAgent string

	mediaType             string
	charset               string
	disallowUnknownFields bool

	// Optional function called after every successful request made to the DO Clients
	beforeRequestDo    BeforeRequestDoCallback
	onRequestCompleted RequestCompletionCallback
}

type BeforeRequestDoCallback func(*http.Client, *http.Request, interface{})

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

func (c *Client) SetHTTPClient(client *http.Client) {
	c.http = client
}

func (c Client) Debug() bool {
	return c.debug
}

func (c *Client) SetDebug(debug bool) {
	c.debug = debug
}

func (c Client) CompanyID() string {
	return c.companyID
}

func (c *Client) SetCompanyID(companyID string) {
	c.companyID = companyID
}

func (c Client) UseTokenAuth() bool {
	return c.useTokenAuth
}

func (c *Client) SetUseTokenAuth(useTokenAuth bool) {
	c.useTokenAuth = useTokenAuth
}

func (c Client) ClientID() string {
	return c.clientID
}

func (c *Client) SetClientID(clientID string) {
	c.clientID = clientID
}

func (c Client) ClientSecret() string {
	return c.clientSecret
}

func (c *Client) SetClientSecret(clientSecret string) {
	c.clientSecret = clientSecret
}

func (c Client) TokenID() string {
	return c.tokenID
}

func (c *Client) SetTokenID(tokenID string) {
	c.tokenID = tokenID
}

func (c Client) TokenSecret() string {
	return c.tokenSecret
}

func (c *Client) SetTokenSecret(tokenSecret string) {
	c.tokenSecret = tokenSecret
}

// func (c Client) AccountID() string {
// 	return c.accountID
// }

// func (c *Client) SetAccountID(accountID string) {
// 	c.accountID = accountID
// }

func (c Client) ContentLanguage() string {
	return c.contentLanguage
}

func (c *Client) SetContentLanguage(contentLanguage string) {
	c.contentLanguage = contentLanguage
}

func (c Client) BaseURL() (*url.URL, error) {
	tmpl, err := template.New("host").Parse(c.baseURL)
	if err != nil {
		return &url.URL{}, err
	}
	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, map[string]interface{}{"account_id": c.companyID})
	if err != nil {
		return &url.URL{}, err
	}
	return url.Parse(buf.String())
}

func (c *Client) SetBaseURL(baseURL string) {
	c.baseURL = baseURL
}

func (c *Client) SetMediaType(mediaType string) {
	c.mediaType = mediaType
}

func (c Client) MediaType() string {
	return mediaType
}

func (c *Client) SetCharset(charset string) {
	c.charset = charset
}

func (c Client) Charset() string {
	return charset
}

func (c *Client) SetUserAgent(userAgent string) {
	c.userAgent = userAgent
}

func (c Client) UserAgent() string {
	return userAgent
}

func (c *Client) SetDisallowUnknownFields(disallowUnknownFields bool) {
	c.disallowUnknownFields = disallowUnknownFields
}

func (c *Client) SetBeforeRequestDo(fun BeforeRequestDoCallback) {
	c.beforeRequestDo = fun
}

func (c *Client) GetEndpointURL(p string, pathParams PathParams) (url.URL, error) {
	clientURL, err := c.BaseURL()
	if err != nil {
		return url.URL{}, err
	}

	parsed, err := url.Parse(p)
	if err != nil {
		log.Fatal(err)
	}
	q := clientURL.Query()
	for k, vv := range parsed.Query() {
		for _, v := range vv {
			q.Add(k, v)
		}
	}
	clientURL.RawQuery = q.Encode()

	clientURL.Path = path.Join(clientURL.Path, parsed.Path)

	tmpl, err := template.New("path").Parse(clientURL.Path)
	if err != nil {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)
	params := pathParams.Params()
	// params["administration_id"] = c.Administration()
	err = tmpl.Execute(buf, params)
	if err != nil {
		return url.URL{}, err
	}

	clientURL.Path = buf.String()
	return *clientURL, nil
}

func (c *Client) NewRequest(ctx context.Context, req Request) (*http.Request, error) {
	// convert body struct to json
	buf := new(bytes.Buffer)
	if req.RequestBodyInterface() != nil {
		err := json.NewEncoder(buf).Encode(req.RequestBodyInterface())
		if err != nil {
			return nil, err
		}
	}

	// create new http request
	u, err := req.URL()
	if err != nil {
		return nil, err
	}

	r, err := http.NewRequest(req.Method(), u.String(), buf)
	if err != nil {
		return nil, err
	}

	// values := url.Values{}
	// err = utils.AddURLValuesToRequest(values, req, true)
	// if err != nil {
	// 	return nil, err
	// }

	// optionally pass along context
	if ctx != nil {
		r = r.WithContext(ctx)
	}

	// set other headers
	r.Header.Add("Content-Type", fmt.Sprintf("%s; charset=%s", c.MediaType(), c.Charset()))
	r.Header.Add("Accept", c.MediaType())
	r.Header.Add("User-Agent", c.UserAgent())

	if c.ContentLanguage() != "" {
		r.Header.Add("Accept-Language", c.ContentLanguage())
		r.Header.Add("Content-Language", c.ContentLanguage())
	}

	if c.UseTokenAuth() {
		headerValue, err := c.TokenBasedAuthorizationHeader(r)
		log.Println("header:", headerValue)
		if err != nil {
			return r, errors.WithStack(err)
		}
		r.Header.Add("Authorization", headerValue)
	}

	return r, nil
}

func (c *Client) TokenBasedAuthorizationHeader(r *http.Request) (string, error) {
	g := c.NewSignatureGenerator(r)
	signature, err := g.Generate()
	if err != nil {
		return "", err
	}

	// OAuth realm="6239966_SB1",
	// oauth_consumer_key="f09493ec1004a7a790738bd2cce1e20686910fd7cbb06687e0d67f9b3ce5d7d3",
	// oauth_token="430b4153e2b43b979ef93ee0be33d89c2a58024743007d045b588d8b1abbf78a",
	// oauth_signature_method="HMAC-SHA256",
	// oauth_timestamp="1668528724",
	// oauth_nonce="85891668386",
	// oauth_version="1.0",
	// oauth_signature="Ds2fEBHxyJvL6OXavV8rrHaMJUyDZi%2By0%2Bu8b%2B1%2Fik8%3D"
	// OAuth realm="6239966_SB1", oauth_token="430b4153e2b43b979ef93ee0be33d89c2a58024743007d045b588d8b1abbf78a", oauth_consumer_key="f09493ec1004a7a790738bd2cce1e20686910fd7cbb06687e0d67f9b3ce5d7d3", oauth_nonce="85891668386", oauth_timestamp="1668528724", oauth_signature_method="HMAC-SHA256", oauth_version="1.0", oauth_signature="Ds2fEBHxyJvL6OXavV8rrHaMJUyDZi%2By0%2Bu8b%2B1%2Fik8%3D"

	return strings.Replace(fmt.Sprintf(`OAuth realm="%s",
oauth_consumer_key="%s",
oauth_token="%s",
oauth_signature_method="%s",
oauth_timestamp="%s",
oauth_nonce="%s",
oauth_version="%s",
oauth_signature="%s"`,
			g.AccountID,                    // realm
			g.ClientID,                     // oauth_consumer key
			g.TokenID,                      // oauth_token
			g.SignatureMethod.String(),     // oauth_signature_method
			strconv.Itoa(int(g.Timestamp)), // timestamp
			g.Nonce,                        // oauth_nonce
			g.Version,                      // oauth_version
			url.QueryEscape(signature),     // oauth_signature
		), "\n", "", -1),
		nil
}

// Do sends an Client request and returns the Client response. The Client response is json decoded and stored in the value
// pointed to by v, or returned as an error if an Client error has occurred. If v implements the io.Writer interface,
// the raw response will be written to v, without attempting to decode it.
func (c *Client) Do(req *http.Request, body interface{}) (*http.Response, error) {
	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		return nil, err
	}

	if c.onRequestCompleted != nil {
		c.onRequestCompleted(req, httpResp)
	}

	// close body io.Reader
	defer func() {
		if rerr := httpResp.Body.Close(); err == nil {
			err = rerr
		}
	}()

	if c.debug == true {
		dump, _ := httputil.DumpResponse(httpResp, true)
		log.Println(string(dump))
	}

	// check if the response isn't an error
	err = CheckResponse(httpResp)
	if err != nil {
		return httpResp, err
	}

	// check the provided interface parameter
	if httpResp == nil {
		return httpResp, nil
	}

	if body == nil {
		return httpResp, err
	}

	if httpResp.ContentLength == 0 {
		return httpResp, nil
	}

	errResp := &ErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, body, errResp)
	if err != nil {
		return httpResp, err
	}

	if errResp.Error() != "" {
		return httpResp, errResp
	}

	return httpResp, nil
}

func (c *Client) Unmarshal(r io.Reader, vv ...interface{}) error {
	if len(vv) == 0 {
		return nil
	}

	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	errs := []error{}
	for _, v := range vv {
		r := bytes.NewReader(b)
		dec := json.NewDecoder(r)
		if c.disallowUnknownFields {
			dec.DisallowUnknownFields()
		}

		err := dec.Decode(v)
		if err != nil && err != io.EOF {
			errs = append(errs, err)
		}

	}

	if len(errs) == len(vv) {
		// Everything errored
		msgs := make([]string, len(errs))
		for i, e := range errs {
			msgs[i] = fmt.Sprint(e)
		}
		return errors.New(strings.Join(msgs, ", "))
	}

	return nil
}

// CheckResponse checks the Client response for errors, and returns them if
// present. A response is considered an error if it has a status code outside
// the 200 range. Client error responses are expected to have either no response
// body, or a json response body that maps to ErrorResponse. Any other response
// body will be silently ignored.
func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}

	// Don't check content-lenght: a created response, for example, has no body
	// if r.Header.Get("Content-Length") == "0" {
	// 	errorResponse.Errors.Message = r.Status
	// 	return errorResponse
	// }

	if c := r.StatusCode; c >= 200 && c <= 299 {
		return nil
	}

	// read data and copy it back
	data, err := ioutil.ReadAll(r.Body)
	r.Body = ioutil.NopCloser(bytes.NewReader(data))
	if err != nil {
		return errorResponse
	}

	err = checkContentType(r)
	if err != nil {
		return errors.WithStack(err)
	}

	if r.ContentLength == 0 {
		return errors.New("response body is empty")
	}

	// convert json to struct
	if len(data) != 0 {
		err = json.Unmarshal(data, &errorResponse)
		if err != nil {
			return errors.WithStack(err)
		}
	}

	if errorResponse.Error() != "" {
		return errorResponse
	}

	return nil
}

// {
//   "type": "https://www.w3.org/Protocols/rfc2616/rfc2616-sec10.html#sec10.4.4",
//   "title": "Forbidden",
//   "status": 403,
//   "o:errorDetails": [
//     {
//       "detail": "The account record is only available as a beta record. Enable the REST Record Service (Beta) feature in Setup > Company > Enable Features to work with this record.",
//       "o:errorCode": "INSUFFICIENT_PERMISSION"
//     }
//   ]
// }

type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response

	Type         string       `json:"type"`
	Title        interface{}  `json:"title"`
	Status       int          `json:"status"`
	ErrorDetails ErrorDetails `json:"o:errorDetails"`
}

func (r *ErrorResponse) Error() string {
	errors := []string{}

	for _, d := range r.ErrorDetails {
		err := d.Error()
		if err != "" {
			errors = append(errors, err)
		}
	}

	return strings.Join(errors, "\r\n")
}

type ErrorDetails []ErrorDetail

type ErrorDetail struct {
	Detail    string `json:"detail"`
	ErrorCode string `json:"o:errorCode"`
}

func (d *ErrorDetail) Error() string {
	if d.ErrorCode != "" {
		return fmt.Sprintf("%s: %s", d.ErrorCode, d.Detail)
	}
	return ""
}

func checkContentType(response *http.Response) error {
	header := response.Header.Get("Content-Type")
	contentType := strings.Split(header, ";")[0]
	if contentType != "application/vnd.oracle.resource+json" {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

func (c *Client) NewSignatureGenerator(r *http.Request) *SignatureGenerator {
	return &SignatureGenerator{
		SignatureMethod:   HMACSHA256,
		BaseURL:           r.URL.String(),
		HTTPRequestMethod: r.Method,
		ClientID:          c.ClientID(),
		ClientSecret:      c.ClientSecret(),
		TokenID:           c.TokenID(),
		TokenSecret:       c.TokenSecret(),
		AccountID:         strings.Replace(c.CompanyID(), "-", "_", -1),
		Nonce:             GenerateNonce(),
		// Nonce:   "35878027199",
		Version:   "1.0",
		Timestamp: time.Now().Unix(),
		// Timestamp: 1668529334,
	}
	// return &SignatureGenerator{
	// 	SignatureMethod:   HMACSHA256,
	// 	BaseURL:           "https://1234567.restlets.api.netsuite.com/rest/accesstoken",
	// 	HTTPRequestMethod: "POST",
	// 	ClientID:          "60712990bc09623786e7047c226bcb3f86d49dca0b04efc21001dc76d97a81f5",
	// 	ClientSecret:      "60712990bc09623786e7047c226bcb3f86d49dca0b04efc21001dc76d97a81f5",
	// 	TokenID:           "447d0cba5569a2d616e32a537110bc8c10ebcf42cc1fa34d6f76d08531abc179",
	// 	TokenSecret:       "447d0cba5569a2d616e32a537110bc8c10ebcf42cc1fa34d6f76d08531abc179",
	// 	AccountID:         "1234567",
	// 	Nonce:             "wjRgXQPWhYtKl0A7bO8Z",
	// 	Version:           "1.0",
	// 	Timestamp:         1576079512,
	// }
}
