package data

type Course struct {
	Id       uint   `json:"id"`
	Name     string `json:"course_name"`
	Duration uint   `json:"duration"`
	UserId   uint   `json:"user_id,omitempty"`
}
