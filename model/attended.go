package model

type AttendedActivity struct {
	ActivityID uint32 `json:"activity_id"`
	StudentID  string `json:"student_id"`
	TxID       []byte `json:"transaction_id"`
}

func NewAttendedActivity(activityID uint32, studentID string, txid []byte) *AttendedActivity {
	return &AttendedActivity{
		ActivityID: activityID,
		StudentID:  studentID,
		TxID:       txid,
	}
}
