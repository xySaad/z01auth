package z01auth

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/xySaad/z01auth/graphql/gqlgen"
	"golang.org/x/oauth2"
)

var ErrMultipleUsers = fmt.Errorf("Found multiple users with the same gitea id")

func determineRole(user gqlgen.GetPublicUserPublicUserUser_public_view) CandidateRole {
	if len(user.Module) > 0 {
		return CandidateRole_TALENT
	}
	if len(user.PiscineGo) > 0 {
		return CandidateRole_POOLER
	}

	return CandidateRole_NONE
}

func (c *Config) Callback(code string) (*Candidate, error) {
	config := c.config
	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return nil, err
	}

	client := config.Client(context.Background(), token)
	resp, err := client.Get(giteaEndpoint("/api/v1/user"))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	giteaUser := &GiteaUser{}
	err = json.NewDecoder(resp.Body).Decode(giteaUser)
	if err != nil {
		return nil, err
	}

	publicUser, err := gqlgen.GetPublicUser(context.Background(), c.graphqlClient, giteaUser.ID)
	if err != nil {
		return nil, err
	}
	json.NewEncoder(os.Stdout).Encode(publicUser)
	if len(publicUser.PublicUser) > 1 {
		return nil, ErrMultipleUsers
	}
	user := publicUser.PublicUser[0]

	candidate := &Candidate{
		GiteaID:        giteaUser.ID,
		AvatarURL:      giteaUser.AvatarURL,
		Description:    giteaUser.Description,
		GiteaLogin:     giteaUser.Login,
		Role:           determineRole(user),
		GraphqlLogin:   user.Login,
		PlatformAccess: user.CanAccessPlatform,
		Campus:         user.Campus,
		GraphqlId:      user.Id,
	}
	return candidate, nil
}

func (c *Config) AuthCodeURL(state string, opts ...oauth2.AuthCodeOption) string {
	return c.config.AuthCodeURL(state, opts...)
}
