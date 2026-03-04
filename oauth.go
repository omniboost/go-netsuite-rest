package netsuite

import (
	"context"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const (
	scope                = "*"
	oauthStateString     = ""
	authorizationTimeout = 60 * time.Second
	tokenTimeout         = 5 * time.Second
)

type Oauth2Config struct {
	CompanyID string
	oauth2.Config
}

func NewOauth2Config(companyID string) *Oauth2Config {
	config := &Oauth2Config{
		CompanyID: companyID,
		Config: oauth2.Config{
			RedirectURL:  "",
			ClientID:     "",
			ClientSecret: "",
			Scopes:       []string{scope},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://{{.account_id}}.suitetalk.api.netsuite.com/services/rest/auth/oauth2/v1/token",
			},
		},
	}

	return config
}

func (c *Oauth2Config) Client(ctx context.Context, t *oauth2.Token) *http.Client {
	return c.Config.Client(ctx, t)
}
