package workable

import (
	"net/http"
	"net/url"
)

var _ OAuthService = &oauthService{}

type OAuthService interface {
	CreateAuthURL(AuthorizeURLInput) (string, error)
	GetAccessToken(AccessTokenInput) (accessToken AccessTokenOutput, err error)
	RefreshAccessToken(RefreshTokenInput) (accessToken AccessTokenOutput, err error)
}

type oauthService struct {
	client       *http.Client
	clientID     string
	clientSecret string
}

func NewOAuthService(clientID, clientSecret string, httpClient *http.Client) OAuthService {
	if httpClient == nil {
		httpClient = http.DefaultClient
	}
	return &oauthService{
		client:       httpClient,
		clientID:     clientID,
		clientSecret: clientSecret,
	}
}

func (o *oauthService) CreateAuthURL(d AuthorizeURLInput) (result string, err error) {
	if d.ClientID == "" && o.clientID == "" {
		return "", ErrClientIDMissing
	}
	if d.ClientID == "" {
		d.ClientID = o.clientID
	}
	authURL, err := url.Parse(authorizeURL)
	if err != nil {
		return "", err
	}
	q := authURL.Query()
	q.Add("client_id", d.ClientID)
	if d.RedirectURI != "" {
		q.Add("redirect_uri", d.RedirectURI)
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

func (o *oauthService) GetAccessToken(d AccessTokenInput) (accessToken AccessTokenOutput, err error) {
	panic("not implemented")
}

func (o *oauthService) RefreshAccessToken(RefreshTokenInput) (accessToken AccessTokenOutput, err error) {
	panic("not implemented")
}
