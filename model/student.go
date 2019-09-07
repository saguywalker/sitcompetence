package model

// Student defines student information
type Student struct {
	StudentID  string `json:"student_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Department string `json:"department"`
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
