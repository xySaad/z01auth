package internal

import (
	"net/http"

	"github.com/xySaad/z01auth/types"
)

type HTTPClient struct {
	http.Client
}

func NewHTTPClient(token types.TokenSupplier) HTTPClient {
	transport := NewAuthedTransport(token)
	return HTTPClient{
		Client: http.Client{Transport: &transport},
	}
}
