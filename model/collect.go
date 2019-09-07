package model

// CollectedBadges defines collected badges for corresponding student
type CollectedBadges struct {
	StudentID    string `json:"student_id"`
	CompetenceID uint16 `json:"competence_id"`
	TxID         []byte `json:"tx_id"`
}

// NewCollectedBadges return new CollectedBadges struct
func NewCollectedBadges(studentID string, competenceID uint16, txid []byte) *CollectedBadges {
	return &CollectedBadges{
		StudentID:    studentID,
		CompetenceID: competenceID,
		TxID:         txid,
	}
}
