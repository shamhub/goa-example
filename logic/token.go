package mathapi

import (
	"context"

	token "github.com/shamhub/goa-example/gen/token"
	"goa.design/clue/log"
)

// token service example implementation.
// The example methods log the requests and return zero values.
type tokensrvc struct{}

// NewToken returns the token service implementation.
func NewToken() token.Service {
	return &tokensrvc{}
}

// Accepts username and password in the body and returns JWT OAUTH2/OIDC token
// with the username as a subject, expiring in 1 hour.
// The username and password are not verified, but cannot be empty strings.
func (s *tokensrvc) Auth(ctx context.Context, p *token.User) (res string, err error) {
	log.Printf(ctx, "token.auth")
	return
}
