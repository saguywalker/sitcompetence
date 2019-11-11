package model

// EditProfile contains additional information for each student
type EditProfile struct {
	ProfilePicture string `json:"profile_picture,omitempty"`
	Motto          string `json:"motto"`
}
