module github.com/omniboost/go-netsuite-rest

go 1.26

require (
	github.com/cydev/zero v0.0.0-20160322155811-4a4535dd56e7
	github.com/golang-jwt/jwt/v5 v5.3.1
	github.com/gorilla/schema v0.0.0-20171211162101-9fa3b6af65dc
	github.com/omniboost/go-httperr v0.0.0-20251103155253-030b17131c87
	github.com/omniboost/oauth-jwtassertion v0.0.0-20260305145153-971d0bf7341a
	github.com/pkg/errors v0.9.1
	gitlab.com/tozd/go/errors v0.11.1
	golang.org/x/oauth2 v0.35.0
	gopkg.in/guregu/null.v3 v3.5.0
)

replace github.com/gorilla/schema => github.com/omniboost/schema v1.1.1-0.20211111150515-2e872025e306
