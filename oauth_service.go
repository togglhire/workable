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
	info   OAuthServiceInput
	client *Client
}

func (s *oauthServiceImpl) CreateAuthURL(d AuthorizeURLInput) (result string, err error) {
	if s.info.ClientID == "" {
		return "", ErrClientIDMissing
	}

	authorizeURL := strings.Replace(defaultBaseURL+"/oauth/authorize", "{subdomain}", "www", -1)
	authorizeURL = strings.Replace(authorizeURL, "{domain}", s.client.domain, -1)
	authURL, err := url.Parse(authorizeURL)
	if err != nil {
		return "", err
	}
	q := authURL.Query()
	q.Add("client_id", s.info.ClientID)
	if s.info.RedirectURI != "" {
		q.Add("redirect_uri", s.info.RedirectURI)
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

// GetAccessToken retrieves the access token
func (s *oauthServiceImpl) GetAccessToken(d AccessTokenInput) (token Token, err error) {
	if s.client == nil {
		return token, ErrClientIsNil
	}

	form := url.Values{}
	form.Add("grant_type", string(grantTypeAuthorizationCode))
	form.Add("client_id", s.info.ClientID)
	form.Add("client_secret", s.info.ClientSecret)
	form.Add("redirect_uri", s.info.RedirectURI)
	form.Add("code", d.Code)

	tokenURL := strings.Replace(defaultBaseURL+"/oauth/token", "{subdomain}", "www", -1)
	tokenURL = strings.Replace(tokenURL, "{domain}", s.client.domain, -1)
	accessTokenURL, err := url.Parse(tokenURL)
	if err != nil {
		return token, err
	}

	req, err := http.NewRequest("POST", accessTokenURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return token, err
	}
	req.Header.Set("Accept", "application/json")

	err = do(s.client.httpClient, req, &token)
	if err != nil {
		return token, err
	}

	return token, err
}

// RefreshAccessToken retrieves a new access token
func (s *oauthServiceImpl) RefreshAccessToken(d RefreshTokenInput) (token Token, err error) {
	if s.client == nil {
		return token, ErrClientIsNil
	}

	form := url.Values{}
	form.Add("grant_type", string(grantTypeRefreshToken))
	form.Add("client_id", s.info.ClientID)
	form.Add("client_secret", s.info.ClientSecret)

	if d.RefreshToken != "" {
		if s.client.token.RefreshToken == "" {
			d.RefreshToken = s.client.token.RefreshToken
		}
	}

	if d.RefreshToken == "" {
		return token, ErrRefreshTokenMissing
	}

	form.Add("refresh_token", d.RefreshToken)

	tokenURL := strings.Replace(defaultBaseURL+"/oauth/token", "{subdomain}", "www", -1)
	tokenURL = strings.Replace(tokenURL, "{domain}", s.client.domain, -1)
	accessTokenURL, err := url.Parse(tokenURL)
	if err != nil {
		return token, err
	}

	req, err := http.NewRequest("POST", accessTokenURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return token, err
	}
	req.Header.Set("Accept", "application/json")

	err = do(s.client.httpClient, req, &token)
	if err != nil {
		return token, err
	}

	return token, err
}

func (s *oauthServiceImpl) RevokePermissions() (err error) {
	if s.client == nil {
		return ErrClientIsNil
	}

	if s.client.token.AccessToken == "" {
		return ErrAccessTokenMissing
	}

	if s.client.token.RefreshToken == "" {
		return ErrRefreshTokenMissing
	}

	form := url.Values{}
	form.Add("refresh_token", s.client.token.RefreshToken)

	revokeURLTemplate := strings.Replace(defaultBaseURL+"/oauth/revoke", "{subdomain}", "www", -1)
	revokeURLTemplate = strings.Replace(revokeURLTemplate, "{domain}", s.client.domain, -1)
	revokeURL, err := url.Parse(revokeURLTemplate)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", revokeURL.String(), strings.NewReader(form.Encode()))
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/json")

	dummyStruct := struct{}{}
	err = do(s.client.httpClient, req, &dummyStruct)
	return err
}
