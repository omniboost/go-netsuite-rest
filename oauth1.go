package netsuite

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"hash"
	"log"
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
	requestParameters := []string{
		fmt.Sprintf("oauth_consumer_key=%s", g.ClientID),
		fmt.Sprintf("oauth_nonce=%s", g.Nonce),
		fmt.Sprintf("oauth_signature_method=%s", g.SignatureMethod.String()),
		fmt.Sprintf("oauth_timestamp=%s", strconv.Itoa(int(g.Timestamp))),
		fmt.Sprintf("oauth_token=%s", g.TokenID),
		fmt.Sprintf("oauth_version=%s", g.Version),
	}

	for _, v := range requestParameters {
		log.Println(v)
	}

	// for i, v := range requestParameters {
	// 	requestParameters[i] = url.QueryEscape(v)
	// }
	normalizedRequestParameters := strings.Join(requestParameters, "&")

	dataPieces := []string{
		g.HTTPRequestMethod,         // http-request-method
		g.BaseURL,                   // base-string-uri
		normalizedRequestParameters, // normalized-request-parameters
	}

	for i, v := range dataPieces {
		dataPieces[i] = url.QueryEscape(v)
	}

	data := strings.Join(dataPieces, "&")
	for _, v := range dataPieces {
		log.Println(v)
	}
	log.Println("data:", data)
	key := g.ClientSecret + "&" + g.TokenSecret
	log.Println("key:", key)

	var (
		signature string
		err       error
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
