//go:generate go run github.com/Khan/genqlient graphql/genqlient.yaml
package z01auth

import (
	"net/http"

	"github.com/Khan/genqlient/graphql"
	"golang.org/x/oauth2"
)

const ORIGIN = "https://learn.zone01oujda.ma"
const GITEA_ENDPOINT = ORIGIN + "/git"
const GRAPHQL_ENDPOINT = ORIGIN + "/api/graphql-engine/v1/graphql"

func giteaEndpoint(path string) string {
	return GITEA_ENDPOINT + path
}

type authedTransport struct {
	token   string
	wrapped http.RoundTripper
}

func (t *authedTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Authorization", "Bearer "+t.token)
	return t.wrapped.RoundTrip(req)
}

type Config struct {
	config        oauth2.Config
	graphqlClient graphql.Client
}

func New(clientID, clientSecret, graphQLToken string) *Config {
	httpClient := &http.Client{
		Transport: &authedTransport{
			token:   graphQLToken,
			wrapped: http.DefaultTransport,
		},
	}
	client := graphql.NewClient(GRAPHQL_ENDPOINT, httpClient)
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:5051/api/auth/callback",
		Scopes:       []string{"read:user"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  giteaEndpoint("/login/oauth/authorize"),
			TokenURL: giteaEndpoint("/login/oauth/access_token")},
	}
	return &Config{config, client}
}
