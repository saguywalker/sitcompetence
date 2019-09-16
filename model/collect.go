package model

// CollectedCompetence defines collected badges for corresponding student
type CollectedCompetence struct {
	StudentID    string `json:"student_id"`
	CompetenceID uint16 `json:"competence_id"`
	TxID         string `json:"tx_id"`
}

// NewCollectedCompetence return new CollectedBadges struct
func NewCollectedCompetence(studentID string, competenceID uint16, txid string) *CollectedCompetence {
	return &CollectedCompetence{
		StudentID:    studentID,
		CompetenceID: competenceID,
		TxID:         txid,
	}
}
