package workable

type CandidateInput struct {
	Sourced   bool      `json:"sourced"`
	Candidate Candidate `json:"candidate"`
}

type Candidate struct {
	ID                     string            `json:"id"`
	Name                   string            `json:"name"`
	Firstname              string            `json:"firstname"`
	Lastname               string            `json:"lastname"`
	Headline               string            `json:"headline"`
	Summary                string            `json:"summary"`
	Address                string            `json:"address"`
	Phone                  string            `json:"phone"`
	Email                  string            `json:"email"`
	CoverLetter            string            `json:"cover_letter"`
	EducationEntries       []EducationEntry  `json:"education_entries"`
	ExperienceEntries      []ExperienceEntry `json:"experience_entries"`
	Skills                 []string          `json:"skills"`
	SocialProfiles         []SocialProfile   `json:"social_profiles"`
	ResumeURL              string            `json:"resume_url"`
	Tags                   []string          `json:"tags"`
	Disqualifed            bool              `json:"disqualifed"`
	DisqualificationReason string            `json:"disqualification_reason"`
	DisqualifiedAt         string            `json:"disqualified_at"`
	Domain                 string            `json:"domain"`
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
