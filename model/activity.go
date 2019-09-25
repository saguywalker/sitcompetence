package model

// Activity table in Postgres
type Activity struct {
	ActivityID   uint32       `json:"activity_id"`
	ActivityName string       `json:"activity_name"`
	Description  string       `json:"description"`
	Date         string       `json:"activity_date"`
	Time         string       `json:"time"`
	Creator      string       `json:"creator"`
	Organizer    string       `json:"organizer"`
	Competences  []Competence `json:"competences"`
	Category     string       `json:"category"`
	Location     string       `json:"location"`
	StudentSite  bool         `json:"student_stie"`
}

// NewActivity creates new activity struct
func NewActivity(activityID uint32, activityName, description, date, time, creator, organizer string, competences []Competence, category, location string, studentSite bool) *Activity {
	return &Activity{
		ActivityID:   activityID,
		ActivityName: activityName,
		Description:  description,
		Date:         date,
		Time:         time,
		Creator:      creator,
		Organizer:    organizer,
		Competences:  competences,
		Category:     category,
		Location:     location,
		StudentSite:  studentSite,
	}
}
