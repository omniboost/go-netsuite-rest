package netsuite

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"hash"
	"math/rand"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// nonce returns a unique string.
func GenerateNonce() string {
	const allowed = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 20)
	for i := range b {
		b[i] = allowed[rand.Intn(len(allowed))]
	}
	return string(b)
}

// SignatureMethod identifies a signature method.
type SignatureMethod int

func (sm SignatureMethod) String() string {
	switch sm {
	case RSASHA1:
		return "RSA-SHA1"
	case RSASHA256:
		return "RSA-SHA256"
	case HMACSHA1:
		return "HMAC-SHA1"
	case HMACSHA256:
		return "HMAC-SHA256"
	case PLAINTEXT:
		return "PLAINTEXT"
	default:
		return "unknown"
	}
}

const (
	HMACSHA1   SignatureMethod = iota // HMAC-SHA1
	RSASHA1                           // RSA-SHA1
	PLAINTEXT                         // Plain text
	HMACSHA256                        // HMAC-256
	RSASHA256                         // RSA-SHA256
)

type SignatureGenerator struct {
	// SignatureMethod specifies the method for signing a request.
	SignatureMethod SignatureMethod

	BaseURL           string
	HTTPRequestMethod string

	ClientID     string // ConsumerKey
	ClientSecret string // ConsumerSecret
	TokenID      string // Token
	TokenSecret  string // TokenSecret
	AccountID    string

	Nonce     string
	Version   string
	Timestamp int64
}

// oauthParams returns the OAuth request parameters for the given credentials,
// method, URL and application parameters. See
// http://tools.ietf.org/html/rfc5849#section-3.4 for more information about
// signatures.
func (g *SignatureGenerator) Generate() (string, error) {
	u, err := url.Parse(g.BaseURL)
	if err != nil {
		return "", err
	}

	baseURL, err := url.Parse(g.BaseURL)
	if err != nil {
		return "", err
	}
	baseURL.RawQuery = ""

	requestParameters := u.Query()
	requestParameters.Add("oauth_consumer_key", g.ClientID)
	requestParameters.Add("oauth_nonce", g.Nonce)
	requestParameters.Add("oauth_signature_method", g.SignatureMethod.String())
	requestParameters.Add("oauth_timestamp", strconv.Itoa(int(g.Timestamp)))
	requestParameters.Add("oauth_token", g.TokenID)
	requestParameters.Add("oauth_version", g.Version)

	// for i, v := range requestParameters {
	// 	requestParameters[i] = url.QueryEscape(v)
	// }

	dataPieces := []string{
		g.HTTPRequestMethod,        // http-request-method
		baseURL.String(),           // base-string-uri
		requestParameters.Encode(), // normalized-request-parameters
	}

	for i, v := range dataPieces {
		dataPieces[i] = url.QueryEscape(v)
	}

	data := strings.Join(dataPieces, "&")
	key := g.ClientSecret + "&" + g.TokenSecret

	var (
		signature string
	)

	switch g.SignatureMethod {
	case HMACSHA1:
		signature, err = g.hmacSignature(sha1.New, data, key)
	case HMACSHA256:
		signature, err = g.hmacSignature(sha256.New, data, key)
	// case RSASHA1:
	// 	signature, err = g.rsaSignature(crypto.SHA1, data, key)
	// case RSASHA256:
	// 	signature, err = g.rsaSignature(crypto.SHA256, data, key)
	case PLAINTEXT:
		signature = g.plainTextSignature(data)
	default:
		err = errors.New("oauth: unknown signature method")
	}

	return signature, err
}

func (g *SignatureGenerator) plainTextSignature(data string) string {
	return data
}

func (g *SignatureGenerator) hmacSignature(h func() hash.Hash, data, key string) (string, error) {
	hm := hmac.New(h, []byte(key))
	_, err := hm.Write([]byte(data))
	return base64.StdEncoding.EncodeToString((hm.Sum(nil))), err
}
