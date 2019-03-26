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

type OAuthServiceInput struct {
	ClientID     string
	ClientSecret string
	RedirectURI  string
}

// AuthorizeURLInput holds the info required to create an Authorize request.
type AuthorizeURLInput struct {
	Scopes []OAuthScope
	State  string
}

// AccessTokenInput holds the info required to retrieve the access token.
type AccessTokenInput struct {
	Code string
}

type RefreshTokenInput struct {
	RefreshToken string
}

type Token struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	CreatedAt    int64  `json:"created_at"` //unix timestamp
}
