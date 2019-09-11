package model

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"

	"github.com/cbergoon/merkletree"
)

// GiveBadge struct will be stored in a blockchain
type GiveBadge struct {
	StudentID    string `json:"student_id"`
	CompetenceID uint16 `json:"competence_id"`
	By           string `json:"by"`
	Semester     uint16 `json:"semester"`
}

// BadgeList contains list of givebadge struct
type BadgeList []GiveBadge

// NewGiveBadge creates new GiveBadge
func NewGiveBadge(studentID string, competenceID uint16, by string, semester uint16) *GiveBadge {
	return &GiveBadge{
		StudentID:    studentID,
		CompetenceID: competenceID,
		By:           by,
		Semester:     semester,
	}
}

// CalculateHash return hash of calling struct
func (g GiveBadge) CalculateHash() ([]byte, error) {
	value, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	if _, err := h.Write(value); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals check hash of 2 structs
func (g GiveBadge) Equals(other merkletree.Content) (bool, error) {
	h1, err := g.CalculateHash()
	if err != nil {
		return false, err
	}

	h2, err := other.CalculateHash()
	if err != nil {
		return false, err
	}

	return bytes.Equal(h1, h2), nil
}

// ApproveActivity will be stored in a blockchain
type ApproveActivity struct {
	StudentID  string `json:"student_id"`
	ActivityID uint32 `json:"activity_id"`
	Approver   string `json:"approver"`
	Semester   uint16 `json:"semester"`
}

// ActivityList contains list of approveactivity struct
type ActivityList []ApproveActivity

// NewApproveActivity creates new ApproveActivity
func NewApproveActivity(studentID string, activityID uint32, approver string, semester uint16) *ApproveActivity {
	return &ApproveActivity{
		StudentID:  studentID,
		ActivityID: activityID,
		Approver:   approver,
		Semester:   semester,
	}
}

// CalculateHash return hash of calling struct
func (a ApproveActivity) CalculateHash() ([]byte, error) {
	value, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}

	h := sha256.New()
	if _, err := h.Write(value); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}

// Equals check hash of 2 structs
func (a ApproveActivity) Equals(other merkletree.Content) (bool, error) {
	h1, err := a.CalculateHash()
	if err != nil {
		return false, err
	}

	h2, err := other.CalculateHash()
	if err != nil {
		return false, err
	}

	return bytes.Equal(h1, h2), nil
}
