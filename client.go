package netsuite

import (
	"bufio"
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
	"unicode"
	"unicode/utf8"

	httperr "github.com/omniboost/go-httperr"
	"gitlab.com/tozd/go/errors"
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

	return r, nil
}

func (c *Client) TokenBasedAuthorizationHeader(r *http.Request) (string, error) {
	g := c.NewSignatureGenerator(r)
	signature, err := g.Generate()
	if err != nil {
		return "", err
	}

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
	if c.UseTokenAuth() {
		headerValue, err := c.TokenBasedAuthorizationHeader(req)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		req.Header.Add("Authorization", headerValue)
	}

	if c.beforeRequestDo != nil {
		c.beforeRequestDo(c.http, req, body)
	}

	if c.debug == true {
		dump, _ := httputil.DumpRequestOut(req, true)
		log.Println(string(dump))
	}

	httpResp, err := c.http.Do(req)
	if err != nil {
		if httpResp != nil {
			// wrap error in http error so we can handle it properly
			return nil, &httperr.Error{StatusCode: httpResp.StatusCode, Err: err}
		}
		return nil, errors.WithStack(err)
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
	err = checkContentType(httpResp)
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

	// read 512 bytes without draining the response body we may perhaps later
	// use for error handling
	peeked, _ := peek(httpResp, 512)

	errResp := &ErrorResponse{Response: httpResp}
	err = c.Unmarshal(httpResp.Body, body, errResp)
	if err != nil {
		// only return error when its not a json syntax error, otherwise we
		// assume the response body is not json and ignore it
		syntaxErr := &json.SyntaxError{}
		if !errors.As(err, &syntaxErr) {
			return httpResp, err
		}
	}

	// we got a response matching the ErrorResponse struct, use it
	if errResp.Error() != "" {
		return httpResp, errResp
	}

	// no matches on th struct, but we got an error status code, try to return
	// the response body as error message if its just plain text
	if httpResp.StatusCode != 0 && (httpResp.StatusCode < 200 || httpResp.StatusCode > 299) {
		// here the original response body could be just text
		if isPlainText(peeked) {
			return httpResp, errors.New(string(peeked))
		}

		// not text, but still an error status code, return the status as error
		// message
		return httpResp, &httperr.Error{StatusCode: httpResp.StatusCode, Err: errors.New(httpResp.Status)}
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
		return errors.Join(errs...)
	}

	return nil
}

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
	if contentType != "application/json" && contentType != "application/vnd.oracle.resource+json" {
		return fmt.Errorf("Expected Content-Type \"%s\", got \"%s\"", mediaType, contentType)
	}

	return nil
}

func (c *Client) NewSignatureGenerator(r *http.Request) *SignatureGenerator {
	// u := r.URL
	// u.RawQuery = ""

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
		Version:           "1.0",
		Timestamp:         time.Now().Unix(),
	}
	// return &SignatureGenerator{
	// 	SignatureMethod:   HMACSHA256,
	// 	BaseURL:           "https://123456.restlets.api.netsuite.com/app/site/hosting/restlet.nl?script=6&deploy=1&customParam=someValue&testParam=someOtherValue",
	// 	HTTPRequestMethod: "POST",
	// 	ClientID:          "ef40afdd8abaac111b13825dd5e5e2ddddb44f86d5a0dd6dcf38c20aae6b67e4",
	// 	ClientSecret:      "d26ad321a4b2f23b0741c8d38392ce01c3e23e109df6c96eac6d099e9ab9e8b5",
	// 	TokenID:           "2b0ce516420110bcbd36b69e99196d1b7f6de3c6234c5afb799b73d87569f5cc",
	// 	TokenSecret:       "c29a677df7d5439a458c063654187e3d678d73aca8e3c9d8bea1478a3eb0d295",
	// 	AccountID:         "123456",
	// 	Nonce:             "fjaLirsIcCGVZWzBX0pg",
	// 	Version:           "1.0",
	// 	Timestamp:         1508242306,
	// }
}

func isPlainText(b []byte) bool {
	s := bytes.TrimSpace(b)
	return utf8.Valid(s) && bytes.IndexFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) &&
			!unicode.IsDigit(r) &&
			!unicode.IsSpace(r) &&
			!strings.ContainsRune(`.,!?;:'"()-_/\@#%&*+=`, r)
	}) == -1
}

func peek(resp *http.Response, n int) ([]byte, error) {
	br := bufio.NewReaderSize(resp.Body, n)
	peeked, err := br.Peek(n)
	if err != nil && err != io.EOF && err != bufio.ErrBufferFull {
		return nil, err
	}

	// Replace the response body with the buffered reader
	resp.Body = io.NopCloser(br)

	return peeked, nil
}
