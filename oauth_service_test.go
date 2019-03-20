package workable

import (
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
