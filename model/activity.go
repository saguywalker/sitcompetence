package model

// Activity table in Postgres
type Activity struct {
	ActivityID    uint32       `json:"activity_id"`
	ActivityName  string       `json:"activity_name"`
	Description   string       `json:"description"`
	Date          string       `json:"activity_date"`
	Time          string       `json:"time"`
	Creator       string       `json:"creator"`
	Organizer     string       `json:"organizer"`
	Category      string       `json:"category"`
	Location      string       `json:"location"`
	Semester      uint16       `json:"semester"`
	CompetencesID []uint16     `json:"competences_id,omitempty"` // For request
	Competences   []Competence `json:"competences,omitempty"`    // For response
	Attendees     []Student    `json:"attendees,omitempty"`
	StudentSite   bool         `json:"student_stie"`
}

// NewActivity creates new activity struct
func NewActivity(activityID uint32, activityName, description, date, time, creator, organizer, category, location string, semester uint16, studentSite bool) *Activity {
	return &Activity{
		ActivityID:   activityID,
		ActivityName: activityName,
		Description:  description,
		Date:         date,
		Time:         time,
		Creator:      creator,
		Organizer:    organizer,
		Category:     category,
		Location:     location,
		Semester:     semester,
		StudentSite:  studentSite,
	}
}
