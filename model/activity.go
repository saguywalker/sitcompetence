package model

// Activity table in Postgres
type Activity struct {
	ActivityID   uint32 `json:"activity_id"`
	ActivityName string `json:"activity_name"`
	Date         string `json:"activity_date"`
	Creator      string `json:"creator"`
	StudentSite  bool   `json:"student_stie"`
}

// NewActivity creates new activity struct
func NewActivity(activityID uint32, activityName, date, creator string, studentSite bool) *Activity {
	return &Activity{
		ActivityID:   activityID,
		ActivityName: activityName,
		Date:         date,
		Creator:      creator,
		StudentSite:  studentSite,
	}
}
