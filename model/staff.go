package model

// Staff defines staff information
type Staff struct {
	StaffID   string `json:"staff_id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	PublicKey string `json:"publickey"`
}

// NewStaff returns new staff struct
func NewStaff(staffID, firstName, lastName string) *Staff {
	return &Staff{
		StaffID:   staffID,
		FirstName: firstName,
		LastName:  lastName,
	}
}
