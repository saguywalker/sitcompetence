package model

// Staff defines staff information
type Staff struct {
	StaffID   string `json:"staff_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

// NewStaff returns new staff struct
func NewStaff(staffID, firstName, lastName string) *Staff {
	return &Staff{
		StaffID:   staffID,
		FirstName: firstName,
		LastName:  lastName,
	}
}
