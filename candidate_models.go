package workable

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
	Skills                 []string          `json:"skills,omitempty"`
	SocialProfiles         []SocialProfile   `json:"social_profiles,omitempty"`
	ResumeURL              string            `json:"resume_url,omitempty"`
	Tags                   []string          `json:"tags,omitempty"`
	Disqualifed            bool              `json:"disqualifed,omitempty"`
	DisqualificationReason string            `json:"disqualification_reason,omitempty"`
	DisqualifiedAt         string            `json:"disqualified_at,omitempty"`
	Domain                 string            `json:"domain,omitempty"`
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
