package model

// GiveBadge struct will be stored in a blockchain
type GiveBadge struct {
	StudentID    string `json:"student_id"`
	CompetenceID uint16 `json:"competence_id"`
	By           string `json:"by"`
	Semester     uint16 `json:"semester"`
}

// NewGiveBadge creates new GiveBadge
func NewGiveBadge(studentID string, competenceID uint16, by string, semester uint16) *GiveBadge {
	return &GiveBadge{
		StudentID:    studentID,
		CompetenceID: competenceID,
		By:           by,
		Semester:     semester,
	}
}

// ApproveActivity will be stored in a blockchain
type ApproveActivity struct {
	StudentID  string `json:"student_id"`
	ActivityID uint32 `json:"activity_id"`
	Approver   string `json:"approver"`
	Semester   uint16 `json:"semester"`
}

// NewApproveActivity creates new ApproveActivity
func NewApproveActivity(studentID string, activityID uint32, approver string, semester uint16) *ApproveActivity {
	return &ApproveActivity{
		StudentID:  studentID,
		ActivityID: activityID,
		Approver:   approver,
		Semester:   semester,
	}
}
