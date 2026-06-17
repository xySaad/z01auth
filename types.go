package z01auth

import "time"

type GiteaUser struct {
	Active            bool      `json:"active"`
	AvatarURL         string    `json:"avatar_url"`
	Created           time.Time `json:"created"`
	Description       string    `json:"description"`
	Email             string    `json:"email"`
	FollowersCount    int       `json:"followers_count"`
	FollowingCount    int       `json:"following_count"`
	FullName          string    `json:"full_name"`
	HTMLURL           string    `json:"html_url"`
	ID                int       `json:"id"`
	IsAdmin           bool      `json:"is_admin"`
	Language          string    `json:"language"`
	LastLogin         time.Time `json:"last_login"`
	Location          string    `json:"location"`
	Login             string    `json:"login"`
	LoginName         string    `json:"login_name"`
	ProhibitLogin     bool      `json:"prohibit_login"`
	Restricted        bool      `json:"restricted"`
	SourceID          int       `json:"source_id"`
	StarredReposCount int       `json:"starred_repos_count"`
	Visibility        string    `json:"visibility"`
	Website           string    `json:"website"`
}

type CandidateRole string

const (
	CandidateRole_NONE   CandidateRole = "none"
	CandidateRole_STAFF  CandidateRole = "staff"
	CandidateRole_TALENT CandidateRole = "talent"
	CandidateRole_POOLER CandidateRole = "pooler"
)

type Candidate struct {
	GiteaID        int
	AvatarURL      string
	Description    string
	GiteaLogin     string
	Role           CandidateRole
	GraphqlLogin   string
	Campus         string
	PlatformAccess bool
}
