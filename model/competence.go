package model

// Competence table in Postgres
type Competence struct {
	CompetenceID            uint16 `json:"competence_id"`
	CompetenceName          string `json:"competence_name"`
	Description             string `json:"description"`
	BadgeIconURL            string `json:"badge_url"`
	TotalActivitiesRequired uint16 `json:"activities_required"`
}

// NewCompetence creates new Competence struct
func NewCompetence(competenceID uint16, competenceName, description, badgeURL string, required uint16) *Competence {
	return &Competence{
		CompetenceID:            competenceID,
		CompetenceName:          competenceName,
		Description:             description,
		BadgeIconURL:            badgeURL,
		TotalActivitiesRequired: required,
	}
}
