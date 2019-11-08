package model

// Student defines student information
type Student struct {
	StudentID   string                `json:"student_id"`
	FirstName   string                `json:"firstname"`
	LastName    string                `json:"lastname"`
	Department  string                `json:"department"`
	Motto       string                `json:"motto,omitempty"`
	ProfilePath string                `json:"profile_path,omitempty"`
	Collected   []CollectedCompetence `json:"collected_competence,omitempty"`
	Evidence    []byte                `json:"evidence,omitempty"`
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
