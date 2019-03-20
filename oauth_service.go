package workable

import (
	"net/http"
	"net/url"
	"strings"
)

var _ oauthService = &oauthServiceImpl{}

type oauthService interface {
	CreateAuthURL(AuthorizeURLInput) (string, error)
	GetAccessToken(AccessTokenInput) (token Token, err error)
	RefreshAccessToken(RefreshTokenInput) (token Token, err error)
	RevokePermissions() (err error)
}

type oauthServiceImpl struct {
	client *Client
}

func (o *oauthServiceImpl) CreateAuthURL(d AuthorizeURLInput) (result string, err error) {
	if o.client == nil {
		return result, ErrClientIsNil
	}

	if o.client.clientID == "" {
		return "", ErrClientIDMissing
	}

	authURL, err := url.Parse(authorizeURL)
	if err != nil {
		return "", err
	}
	q := authURL.Query()
	q.Add("client_id", o.client.clientID)
	if o.client.redirectURI != "" {
		q.Add("redirect_uri", o.client.redirectURI)
	}

	// TODO: check what kind of values resource can have
	q.Add("resource", "user")

	q.Add("response_type", "code")
	if len(d.Scopes) != 0 {
		q.Add("scope", spaceDelimit(d.Scopes))
	}

	// TODO: check if state is available
	// if d.State != "" {
	// 	q.Add("state", d.State)
	// }
	authURL.RawQuery = q.Encode()
	return authURL.String(), nil
}

// GetAccessToken retrieves the access token and updates the client to use the new access token
func (o *oauthServiceImpl) GetAccessToken(d AccessTokenInput) (token Token, err error) {
	if o.client == nil {
		return token, ErrClientIsNil
	}

	form := url.Values{}
	form.Add("grant_type", string(grantTypeAuthorizationCode))
	form.Add("client_id", o.client.clientID)
	form.Add("client_secret", o.client.clientSecret)
	form.Add("redirect_uri", o.client.redirectURI)
	form.Add("code", d.Code)

	accessTokenURL, err := url.Parse(tokenURL)
	if err != nil {
		return token, err
	}

	req, err := http.NewRequest("POST", accessTokenURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return token, err
	}
	req.Header.Set("Accept", "application/json")

	err = do(o.client.httpClient, req, &token)
	if err != nil {
		return token, err
	}

	o.client.SetAccessToken(&token)
	return token, err
}

// RefreshAccessToken retrieves a new access token and updates the client to use the new access token
func (o *oauthServiceImpl) RefreshAccessToken(d RefreshTokenInput) (token Token, err error) {
	if o.client == nil {
		return token, ErrClientIsNil
	}

	form := url.Values{}
	form.Add("grant_type", string(grantTypeRefreshToken))
	form.Add("client_id", o.client.clientID)
	form.Add("client_secret", o.client.clientSecret)

	if d.RefreshToken != "" {
		if o.client.token != nil {
			d.RefreshToken = o.client.token.RefreshToken
		}
	}

	if d.RefreshToken == "" {
		return token, ErrRefreshTokenMissing
	}

	form.Add("refresh_token", d.RefreshToken)

	accessTokenURL, err := url.Parse(tokenURL)
	if err != nil {
		return token, err
	}

	req, err := http.NewRequest("POST", accessTokenURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return token, err
	}
	req.Header.Set("Accept", "application/json")

	err = do(o.client.httpClient, req, &token)
	if err != nil {
		return token, err
	}

	o.client.SetAccessToken(&token)
	return token, err
}

func (o *oauthServiceImpl) RevokePermissions() (err error) {
	if o.client == nil {
		return ErrClientIsNil
	}

	if o.client.token.AccessToken == "" {
		return ErrAccessTokenMissing
	}

	if o.client.token.RefreshToken == "" {
		return ErrRefreshTokenMissing
	}

	form := url.Values{}
	form.Add("refresh_token", o.client.token.RefreshToken)

	revokeURL, err := url.Parse(revokeURL)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", revokeURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	dummyStruct := struct{}{}
	err = do(o.client.httpClient, req, &dummyStruct)
	return err
}
