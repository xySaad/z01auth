package internal

import (
	"net/http"

	"github.com/xySaad/z01auth/types"
)

type AuthedTransport struct {
	token types.TokenSupplier
}

func NewAuthedTransport(token types.TokenSupplier) AuthedTransport {
	return AuthedTransport{token: token}
}

func (at *AuthedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+at.token.Get())
	return http.DefaultTransport.RoundTrip(req)
}
