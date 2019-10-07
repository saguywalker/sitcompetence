package model

// Student defines student information
type Student struct {
	StudentID  string                `json:"student_id"`
	FirstName  string                `json:"firstname"`
	LastName   string                `json:"lastname"`
	Department string                `json:"department"`
	Collected  []CollectedCompetence `json:"collected_competence,omitempty"`
}

// NewStudent returns new Student struct
func NewStudent(studentID, firstName, lastName, department string) *Student {
	return &Student{
		StudentID:  studentID,
		FirstName:  firstName,
		LastName:   lastName,
		Department: department,
	}
}
