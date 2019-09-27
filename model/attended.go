package model

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
)

// AttendedActivity table in postgres
type AttendedActivity struct {
	ActivityID    uint32 `json:"activity_id"`
	StudentID     string `json:"student_id"`
	Approver      string `json:"approver"`
	TransactionID []byte `json:"transaction_id,omitempty"`
}

// NewAttendedActivity return new struct
func NewAttendedActivity(activityID uint32, studentID, approver string, txid []byte) *AttendedActivity {
	return &AttendedActivity{
		ActivityID:    activityID,
		StudentID:     studentID,
		Approver:      approver,
		TransactionID: txid,
	}
}

// CalculateHash return hash of calling struct
func (a AttendedActivity) CalculateHash() ([]byte, error) {
	value, err := json.Marshal(a)
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
func (a AttendedActivity) Equals(other merkletree.Content) (bool, error) {
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
*/