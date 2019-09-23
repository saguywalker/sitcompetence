package model

// Activity table in Postgres
type Activity struct {
	ActivityID   uint32       `json:"activity_id"`
	ActivityName string       `json:"activity_name"`
	Description  string       `json:"description"`
	Date         string       `json:"activity_date"`
	Creator      string       `json:"creator"`
	Competences  []Competence `json:"competences"`
	StudentSite  bool         `json:"student_stie"`
}

// NewActivity creates new activity struct
func NewActivity(activityID uint32, activityName, description, date, creator string, competences []Competence, studentSite bool) *Activity {
	return &Activity{
		ActivityID:   activityID,
		ActivityName: activityName,
		Description:  description,
		Date:         date,
		Creator:      creator,
		Competences:  competences,
		StudentSite:  studentSite,
	}
}
