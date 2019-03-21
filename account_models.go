package workable

type Accounts struct {
	Accounts []Account `json:"accounts"`
}

type Account struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Subdomain   string `json:"subdomain"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
	WebsiteURL  string `json:"website_url"`
}
