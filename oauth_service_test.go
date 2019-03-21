package workable

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_oauthService_CreateAuthURL(t *testing.T) {
	input := AuthorizeURLInput{
		Scopes: []OAuthScope{
			OAuthScopeReadJobs,
			OAuthScopeReadCandidates,
			OAuthScopeWriteCandidates,
		},
	}
	want := "https://www.workable.com/oauth/authorize?client_id=client_id&redirect_uri=redirect_uri&resource=user&response_type=code&scope=r_jobs+r_candidates+w_candidates"

	client := NewClient(Token{}, nil)
	got, err := client.OAuth(OAuthServiceInput{
		ClientID:     "client_id",
		ClientSecret: "client_secret",
		RedirectURI:  "redirect_uri",
	}).CreateAuthURL(input)

	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func Test_oauthService_GetAccessToken(t *testing.T) {
	t.Skip("Requires working client id, secret & token to work")
}
