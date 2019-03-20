package workable

import (
	"log"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_oauthService_CreateAuthURL(t *testing.T) {
	type args struct {
		data AuthorizeURLInput
	}
	test := struct {
		args    args
		wantURL string
		wantErr bool
	}{
		args: args{
			data: AuthorizeURLInput{
				Scopes: []OAuthScope{
					OAuthScopeReadJobs,
					OAuthScopeReadCandidates,
					OAuthScopeWriteCandidates,
				},
			},
		},
		wantURL: "https://www.workable.com/oauth/authorize?client_id=client_id&redirect_uri=redirect_uri&resource=user&response_type=code&scope=r_jobs+r_candidates+w_candidates",
	}

	client := NewClient("client_id", "client_secret", "redirect_uri", nil, nil)
	gotURL, err := client.OAuth.CreateAuthURL(test.args.data)

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	assert.Equal(t, test.wantURL, gotURL)
}

func Test_oauthService_GetAccessToken(t *testing.T) {
	t.Skip("Requires working client id, secret & token to work")

	requestURI, err := url.Parse("http://partner.com/redirect?code=e7fcd407a73dbe5219")
	assert.NoError(t, err)
	code := requestURI.Query().Get("code")
	type args struct {
		data AccessTokenInput
	}
	test := struct {
		args    args
		wantErr bool
	}{
		args: args{
			data: AccessTokenInput{
				Code: code,
			},
		}}

	client := NewClient("client_id", "client_secret", "redirect_uri", nil, nil)
	accessToken, err := client.OAuth.GetAccessToken(test.args.data)

	switch test.wantErr {
	case true:
		assert.Error(t, err)
	case false:
		assert.NoError(t, err)
	}
	log.Printf("accessToken: %#+v\n", accessToken)
}
