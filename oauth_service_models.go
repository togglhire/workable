package workable

type OAuthScope string

const (
	OAuthScopeReadJobs        = OAuthScope("r_jobs")
	OAuthScopeReadCandidates  = OAuthScope("r_candidates")
	OAuthScopeWriteCandidates = OAuthScope("w_candidates")
)

type grantType string

const (
	grantTypeAuthorizationCode = grantType("authorization_code")
	grantTypeRefreshToken      = grantType("refresh_token")
)

// AuthorizeURLInput holds the info required to create an Authorize request.
type AuthorizeURLInput struct {
	ClientID     string
	RedirectURI  string
	Resource     string
	ResponseType string
	Scopes       []OAuthScope
}

// AccessTokenInput holds the info required to retrieve the access token.
type AccessTokenInput struct {
	GrantType    grantType
	ClientID     string
	ClientSecret string
	Code         string
	RedirectURI  string
}

type AccessTokenOutput struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int    `json:"created_at"`
}
