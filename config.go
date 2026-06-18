//go:generate go run github.com/Khan/genqlient graphql/genqlient.yaml
package z01auth

import (
	"github.com/Khan/genqlient/graphql"
	"github.com/xySaad/z01auth/internal"
	"github.com/xySaad/z01auth/types"
	"golang.org/x/oauth2"
)

const ORIGIN = "https://learn.zone01oujda.ma"
const GITEA_ENDPOINT = ORIGIN + "/git"
const GRAPHQL_ENDPOINT = ORIGIN + "/api/graphql-engine/v1/graphql"

func giteaEndpoint(path string) string {
	return GITEA_ENDPOINT + path
}

type Config struct {
	config        oauth2.Config
	graphqlClient graphql.Client
}

func New(clientID, clientSecret, redirectURL string, token types.TokenSupplier) Config {
	httpClient := internal.NewHTTPClient(token)
	client := graphql.NewClient(GRAPHQL_ENDPOINT, &httpClient)
	config := oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirectURL,
		Scopes:       []string{"read:user"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  giteaEndpoint("/login/oauth/authorize"),
			TokenURL: giteaEndpoint("/login/oauth/access_token")},
	}
	return Config{config, client}
}
