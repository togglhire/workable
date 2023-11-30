package workable

import "time"

type CandidateInput struct {
	Sourced   bool      `json:"sourced"`
	Candidate Candidate `json:"candidate"`
}

type Candidate struct {
	ID                     string            `json:"id,omitempty"`
	Name                   string            `json:"name,omitempty"`
	Firstname              string            `json:"firstname,omitempty"`
	Lastname               string            `json:"lastname,omitempty"`
	Headline               string            `json:"headline,omitempty"`
	Summary                string            `json:"summary,omitempty"`
	Address                string            `json:"address,omitempty"`
	Phone                  string            `json:"phone,omitempty"`
	Email                  string            `json:"email,omitempty"`
	CoverLetter            string            `json:"cover_letter,omitempty"`
	EducationEntries       []EducationEntry  `json:"education_entries,omitempty"`
	ExperienceEntries      []ExperienceEntry `json:"experience_entries,omitempty"`
	Skills                 []Skill           `json:"skills,omitempty"`
	SocialProfiles         []SocialProfile   `json:"social_profiles,omitempty"`
	ResumeURL              string            `json:"resume_url,omitempty"`
	Tags                   []string          `json:"tags,omitempty"`
	Disqualifed            bool              `json:"disqualifed,omitempty"`
	DisqualificationReason string            `json:"disqualification_reason,omitempty"`
	DisqualifiedAt         string            `json:"disqualified_at,omitempty"`
	Domain                 string            `json:"domain,omitempty"`
	CreatedAt              time.Time         `json:"created_at,omitempty"`
	UpdatedAt              time.Time         `json:"updated_at,omitempty"`
}

type CandidateListItem struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
	Headline  string `json:"headline,omitempty"`
	Account   struct {
		Subdomain string `json:"subdomain,omitempty"`
		Name      string `json:"name,omitempty"`
	} `json:"account"`
	Job struct {
		Shortcode string `json:"shortcode,omitempty"`
		Title     string `json:"title,omitempty"`
	} `json:"job"`
	Stage                  string     `json:"stage,omitempty"`
	Disqualifed            bool       `json:"disqualifed,omitempty"`
	DisqualificationReason string     `json:"disqualification_reason,omitempty"`
	Sourced                bool       `json:"sourced,omitempty"`
	ProfileURL             string     `json:"profile_url,omitempty"`
	Email                  string     `json:"email,omitempty"`
	Domain                 string     `json:"domain,omitempty"`
	CreatedAt              time.Time  `json:"created_at,omitempty"`
	UpdatedAt              time.Time  `json:"updated_at,omitempty"`
	HiredAt                *time.Time `json:"hired_at,omitempty"`
	Address                string     `json:"address,omitempty"`
	Phone                  string     `json:"phone,omitempty"`
}

type EducationEntry struct {
	Degree       string `json:"degree"`
	School       string `json:"school"`
	FieldOfStudy string `json:"field_of_study"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

type ExperienceEntry struct {
	Title     string `json:"title"`
	Summary   string `json:"summary"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Current   bool   `json:"current"`
	Company   string `json:"company"`
	Industry  string `json:"industry"`
}

type SocialProfile struct {
	Type     string `json:"type"`
	Name     string `json:"name,omitempty"`
	Username string `json:"username,omitempty"`
	URL      string `json:"url"`
}

type CandidateOutput struct {
	Status    string    `json:"status"`
	Candidate Candidate `json:"candidate"`
}

type ListCandidatesInput struct {
	JobShortCode string
	State        string
	Limit        int
	SinceID      string
	MaxID        string
	CreatedAfter int64
	UpdatedAfter int64
}

type Candidates struct {
	Candidates []CandidateListItem `json:"candidates"`
	Paging     Paging              `json:"paging"`
}

type CandidateUpdateInput struct {
	Firstname         *string           `json:"firstname,omitempty"`
	Lastname          *string           `json:"lastname,omitempty"`
	Email             *string           `json:"email,omitempty"`
	Headline          *string           `json:"headline,omitempty"`
	Summary           *string           `json:"summary,omitempty"`
	Address           *string           `json:"address,omitempty"`
	Phone             *string           `json:"phone,omitempty"`
	CoverLetter       *string           `json:"cover_letter,omitempty"`
	ResumeURL         *string           `json:"resume_url,omitempty"`
	ImageURL          *string           `json:"image_url,omitempty"`
	EducationEntries  []EducationEntry  `json:"education_entries,omitempty"`
	ExperienceEntries []ExperienceEntry `json:"experience_entries,omitempty"`
	Skills            []Skill           `json:"skills,omitempty"`
	Tags              []string          `json:"tags,omitempty"`
	SocialProfiles    []SocialProfile   `json:"social_profiles,omitempty"`
}

type Skill struct {
	Name string `json:"name,omitempty"`
}
