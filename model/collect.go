package model

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
)

// GiveBadgeRequest define givebadge request
type GiveBadgeRequest struct {
	Badges     []CollectedCompetence `json:"badges"`
	PrivateKey string                `json:"sk"`
}

// CollectedCompetence defines collected badges for corresponding student
type CollectedCompetence struct {
	StudentID    string `json:"student_id"`
	CompetenceID uint32 `json:"competence_id"`
	Semester     uint32 `json:"semester"`
	Giver        []byte `json:"giver"`
	TxID         []byte `json:"transaction_id,omitempty"`
	// PrivateKey   string `json:"sk,omitempty"`
}

/*
// NewCollectedCompetence return new CollectedBadges struct
func NewCollectedCompetence(studentID string, competenceID, semester uint32, giver, txid []byte) *CollectedCompetence {
	return &CollectedCompetence{
		StudentID:    studentID,
		CompetenceID: competenceID,
		Semester:     semester,
		Giver:        giver,
		TxID:         txid,
	}
}
*/

// CalculateHash return hash of calling struct
func (c CollectedCompetence) CalculateHash() ([]byte, error) {
	value, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	var trimData bytes.Buffer
	if err := json.Compact(&trimData, value); err != nil {
		return nil, err
	}

	var sortedValue map[string]interface{}
	if err := json.Unmarshal(trimData.Bytes(), &sortedValue); err != nil {
		return nil, err
	}

	jsonBytes, err := json.Marshal(sortedValue)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	if _, err := h.Write(jsonBytes); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

/*
// Equals check hash of 2 structs
func (c CollectedCompetence) Equals(other merkletree.Content) (bool, error) {
	h1, err := c.CalculateHash()
	if err != nil {
		return false, err
	}

	h2, err := other.CalculateHash()
	if err != nil {
		return false, err
	}

	return bytes.Equal(h1, h2), nil
}
*/
